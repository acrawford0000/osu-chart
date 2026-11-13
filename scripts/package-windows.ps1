#Requires -Version 7.0
param(
    [Parameter(Mandatory)][string]$ExePath,
    [Parameter(Mandatory)][string]$OutputDirectory,
    [string]$Version = '1.1.0'
)
$ErrorActionPreference = 'Stop'
if ($Version -notmatch '^\d+\.\d+\.\d+$') { throw 'Version must use major.minor.patch format.' }
$repoPath = Split-Path $PSScriptRoot -Parent
$sourceExe = (Resolve-Path -LiteralPath $ExePath).Path
$binaryInfo = [System.Diagnostics.FileVersionInfo]::GetVersionInfo($sourceExe)
if ($binaryInfo.ProductVersion -ne $Version) { throw 'Executable product version does not match the release version.' }
$stream = [System.IO.File]::OpenRead($sourceExe)
try {
    $reader = [System.IO.BinaryReader]::new($stream)
    $stream.Position = 0x3c
    $peOffset = $reader.ReadInt32()
    $stream.Position = $peOffset + 4
    if ($reader.ReadUInt16() -ne 0x8664) { throw 'Expected a Windows AMD64 executable.' }
} finally { $stream.Dispose() }
New-Item -ItemType Directory -Force -Path $OutputDirectory | Out-Null
$outputPath = (Resolve-Path -LiteralPath $OutputDirectory).Path
$assetBase = "osu-chart-v$Version-windows-amd64"
$exeAsset = Join-Path $outputPath ($assetBase + '.exe')
$zipAsset = Join-Path $outputPath ($assetBase + '.zip')
Copy-Item -LiteralPath $sourceExe -Destination $exeAsset -Force
if (Test-Path -LiteralPath $zipAsset) { Remove-Item -LiteralPath $zipAsset }
$archive = [System.IO.Compression.ZipFile]::Open($zipAsset, [System.IO.Compression.ZipArchiveMode]::Create)
try {
    foreach ($entry in @(
        @{ Source = $sourceExe; Name = 'osu-chart.exe' },
        @{ Source = (Join-Path $repoPath 'LICENSE'); Name = 'LICENSE.txt' },
        @{ Source = (Join-Path $repoPath 'docs\WINDOWS-README.txt'); Name = 'README.txt' }
    )) {
        [System.IO.Compression.ZipFileExtensions]::CreateEntryFromFile($archive, $entry.Source, $entry.Name, [System.IO.Compression.CompressionLevel]::Optimal) | Out-Null
    }
} finally { $archive.Dispose() }
$hashes = foreach ($file in @($exeAsset, $zipAsset)) {
    $hash = (Get-FileHash -LiteralPath $file -Algorithm SHA256).Hash.ToLowerInvariant()
    "$hash  $([System.IO.Path]::GetFileName($file))"
}
Set-Content -LiteralPath (Join-Path $outputPath 'SHA256SUMS.txt') -Value $hashes -Encoding utf8NoBOM
Get-Item -LiteralPath $exeAsset, $zipAsset, (Join-Path $outputPath 'SHA256SUMS.txt') | Select-Object Name, Length

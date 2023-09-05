<script lang="ts">
    import { BrowserOpenURL } from "../../wailsjs/runtime"
    import Picture from "../assets/images/ac_pfp.png"
    
    let youtubeUrl = 'https://www.youtube.com/@ACrawford00';
    let githubUrl = 'https://github.com/acrawford0000';
    
    function openExternalLink(event: MouseEvent) {
        event.preventDefault();
        BrowserOpenURL((event.target as HTMLAnchorElement).href);
        }
</script>
    
<div class="card">
<img src={Picture} alt="ACrawford">

<div class="info">
    <h2>Andrew Crawford</h2>
    <h6>ACrawford</h6>
    
    <a href={youtubeUrl} on:click|preventDefault={openExternalLink} title="Visit my YouTube channel">
    <i class="fab fa-youtube"></i>
    </a>
    
    <a href={githubUrl} on:click|preventDefault={openExternalLink} title="Visit my GitHub profile">
    <i class="fab fa-github"></i>
    </a>
</div>
</div>
  
<style>
    .card {
        background-color: #333;
        color: #fff;
        width: 300px;
        height: 450px;
        padding: 20px;
        border-radius: 10px;
        transition: all 0.3s cubic-bezier(.25,.8,.25,1);
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        align-items: center;
    }
    .card:hover {
        box-shadow: 0px 7px 14px rgba(255, 204, 227,0.25), 0px 10px 10px rgba(0,0,0,0.22);
        transform: scale(1.05);
    }
    .card img {
        width: 50%;
        height: 150px;
        padding: 10%;
        border-radius: 50%;
    }
    .card .info a {
        color: #fff;
        text-decoration: none;
        margin-right: 10px;
        transition: color .3s ease;
    }
    .card .info a:hover {
        color: #1e90ff;
    }
</style>
<script>
    import { onMount, onDestroy } from 'svelte';
    let canvas;
    let ctx;
    let triangles = [];
    let raf;
    let colors = ['#FFB6C1', '#FF69B4', '#FF1493', '#DB7093', '#C71585'];
    let intervalId;
    
    function drawTriangle(x, y, size, color) {
        ctx.beginPath();
        ctx.moveTo(x, y);
        ctx.lineTo(x + size / 2, y + size);
        ctx.lineTo(x - size / 2, y + size);
        ctx.closePath();
        ctx.fillStyle = color;
        ctx.fill();
    }
    
    function addTriangle() {
        let x = Math.random() * canvas.width;
        let y = 0;
        let size = (Math.random() * 20) + 10;
        let color = colors[Math.floor(Math.random() * colors.length)];
        let speed = (Math.random() * 3) + 1;
        let rotationSpeed = (Math.random() * 0.5) - 0.25;
        triangles.push({x, y, size, color, speed, rotationSpeed});
    }
    
    function updateTriangles() {
        for (let i = 0; i < triangles.length; i++) {
            triangles[i].y += triangles[i].speed;
            if (triangles[i].y > canvas.height) {
                triangles.splice(i, 1);
                i--;
            }
        }
    }
    
    function draw() {
        ctx.clearRect(0, 0, canvas.width, canvas.height);
        for (let triangle of triangles) {
            ctx.save();
            ctx.translate(triangle.x, triangle.y);
            ctx.rotate(triangle.rotationSpeed);
            drawTriangle(0, 0, triangle.size, triangle.color);
            ctx.restore();
        }
        updateTriangles();
        raf = requestAnimationFrame(draw);
    }
    
    onMount(() => {
        canvas.width = window.innerWidth;
        canvas.height = window.innerHeight;
        ctx = canvas.getContext('2d');
        
        intervalId = setInterval(addTriangle, 100);
        
        raf = requestAnimationFrame(draw);
        
        return () => {
            cancelAnimationFrame(raf);
            clearInterval(intervalId); // Clear the interval when the component is destroyed.
        }
    });

    onDestroy(() => {
        clearInterval(intervalId); // Clear the interval if the component is destroyed without going through onMount.
    });
</script>

<canvas bind:this={canvas}></canvas>

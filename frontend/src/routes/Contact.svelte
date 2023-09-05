<script>
    import { onMount } from 'svelte';
    import ContactCard from "../components/ContactCard.svelte";
    import BG from "../assets/images/bg.jpg";
    
    let x = 0;
    let y = 0;
    let mouseX = 0;
    let mouseY = 0;
    let strength = 0.15;
    
    onMount(() => {
        const moveBackground = () => {
            x += (mouseX - x) * strength;
            y += (mouseY - y) * strength;
            
            document.documentElement.style.setProperty('--x', `${x}px`);
            document.documentElement.style.setProperty('--y', `${y}px`);
            
            requestAnimationFrame(moveBackground);
        }
        
        moveBackground();
    });
    
    const handleMouseMove = (event) => {
        mouseX = event.clientX - window.innerWidth / 2;
        mouseY = event.clientY - window.innerHeight / 2;
    }
</script>

<style>
    :root {
        --x: 0;
        --y: 0;
    }
    
    .parallax {
        height: 100vh;
        width: 100%;
        overflow: hidden;
        position: relative;
    }
    
    .parallax::before {
        content: "";
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background-image: url(../assets/images/bg.jpg);
        background-size: cover;
        background-position: center center;
        transform: scale(1.1) translate3d(calc(var(--x) * -0.05), calc(var(--y) * -0.05), 0);
        filter: brightness(50%);
    }
    
    .contact-card {
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%) translate3d(calc(var(--x) * -0.025), calc(var(--y) * -0.025), 0);
        box-shadow: calc(var(--x) * -0.025) calc(var(--y) * -0.025) 10px rgba(0, 0, 0, 0.2);
    }
</style>

<div class="parallax" on:mousemove|passive={handleMouseMove}>
    <div class="contact-card">
        <ContactCard />
    </div>
</div>

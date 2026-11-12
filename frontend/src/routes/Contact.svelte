<script>
    import { onMount } from 'svelte';
    import ContactCard from "../components/ContactCard.svelte";
    import BG from "../assets/images/bg.jpg";

    let container;
    let background;
    let card;

    onMount(() => {
        const reducedMotion = window.matchMedia('(prefers-reduced-motion: reduce)');
        const finePointer = window.matchMedia('(hover: hover) and (pointer: fine)');
        let bounds = container.getBoundingClientRect();
        let x = 0;
        let y = 0;
        let targetX = 0;
        let targetY = 0;
        let frame = 0;
        let previousTime = 0;

        const canAnimate = () => !document.hidden && !reducedMotion.matches && finePointer.matches;

        // Only transform the two layers; avoid inherited CSS updates and shadow repaints.
        const render = () => {
            background.style.transform = `scale(1.1) translate3d(${x * -0.05}px, ${y * -0.05}px, 0)`;
            card.style.transform = `translate(-50%, -50%) translate3d(${x * -0.025}px, ${y * -0.025}px, 0)`;
        };

        const animate = (time) => {
            frame = 0;
            // Time-based easing keeps the same speed on high-refresh displays.
            const elapsed = previousTime ? Math.min(time - previousTime, 64) : 1000 / 60;
            previousTime = time;
            const easing = 1 - Math.exp(-elapsed / 100);
            x += (targetX - x) * easing;
            y += (targetY - y) * easing;
            const settled = Math.abs(targetX - x) < 0.1 && Math.abs(targetY - y) < 0.1;
            if (settled) {
                x = targetX;
                y = targetY;
            }
            render();
            if (!settled) frame = requestAnimationFrame(animate);
        };

        const schedule = () => {
            if (!frame && canAnimate()) {
                previousTime = 0;
                frame = requestAnimationFrame(animate);
            }
        };

        const measure = () => { bounds = container.getBoundingClientRect(); };
        const move = (event) => {
            if (!canAnimate() || event.pointerType === 'touch') return;
            targetX = event.clientX - bounds.left - bounds.width / 2;
            targetY = event.clientY - bounds.top - bounds.height / 2;
            schedule();
        };
        const recenter = () => {
            targetX = 0;
            targetY = 0;
            if (x || y) schedule();
        };
        const stop = () => {
            cancelAnimationFrame(frame);
            frame = 0;
            previousTime = 0;
            x = y = targetX = targetY = 0;
            render();
        };
        const updateMotion = () => {
            if (!canAnimate()) stop();
        };
        const resizeObserver = new ResizeObserver(() => {
            measure();
            recenter();
        });

        resizeObserver.observe(container);
        container.addEventListener('pointerenter', measure);
        container.addEventListener('pointermove', move, { passive: true });
        container.addEventListener('pointerleave', recenter);
        document.addEventListener('visibilitychange', updateMotion);
        reducedMotion.addEventListener('change', updateMotion);
        finePointer.addEventListener('change', updateMotion);

        return () => {
            cancelAnimationFrame(frame);
            resizeObserver.disconnect();
            container.removeEventListener('pointerenter', measure);
            container.removeEventListener('pointermove', move);
            container.removeEventListener('pointerleave', recenter);
            document.removeEventListener('visibilitychange', updateMotion);
            reducedMotion.removeEventListener('change', updateMotion);
            finePointer.removeEventListener('change', updateMotion);
        };
    });
</script>

<div class="parallax" bind:this={container}>
    <div class="parallax-background" bind:this={background} style:background-image="url('{BG}')" aria-hidden="true"></div>
    <div class="contact-card" bind:this={card}>
        <ContactCard />
    </div>
</div>

<style>
    .parallax {
        height: 100vh;
        width: 100%;
        overflow: hidden;
        position: relative;
    }

    .parallax-background {
        will-change: transform;
        position: absolute;
        inset: 0;
        background-size: cover;
        background-position: center;
        transform: scale(1.1);
    }

    .parallax-background::after {
        content: "";
        position: absolute;
        inset: 0;
        background: rgba(0, 0, 0, 0.5);
    }

    .contact-card {
        will-change: transform;
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        box-shadow: 0 4px 10px rgba(0, 0, 0, 0.2);
    }

    @media (prefers-reduced-motion: reduce), (hover: none), (pointer: coarse) {
        .parallax-background,
        .contact-card {
            will-change: auto;
        }
    }
</style>

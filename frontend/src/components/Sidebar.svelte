<script>
  import { Router, Route, navigate } from 'svelte-routing';
  import { Icon } from '@smui/common';

  let links = [
    {
      icon: 'list',
      text: 'Players',
      path: '/'
    },
    {
      icon: 'settings',
      text: 'Settings',
      path: '/settings'
    },
    {
      icon: 'stacked_line_chart',
      text: 'Chart',
      path: '/chart'
    },
    {
      icon: 'contact_page',
      text: 'Contact',
      path: '/contact',
    }
  ];

  let active = links[0];
  let activeBarY = 0;
  let sidebarItems = [];
  function handleLinkClick(link) {
    active = link;
    navigate(link.path);
  }

  function handleKeyDown(event, link) {
    if (event.key === 'Enter') {
      handleLinkClick(link);
    }
  }
  
  $: {
    const activeIndex = links.findIndex(link => link.path === active.path);
    if (sidebarItems[activeIndex]) {
      activeBarY = sidebarItems[activeIndex].offsetTop;
    }
  }
  </script>

<div class="sidebar">
  <Router>
    <div class="active-bar" style="transform: translateY({activeBarY}px)"></div>
    {#each links as link, index}
      {#if index === links.length - 1}
        <div class="divider"></div>
        <!-- svelte-ignore a11y-no-noninteractive-tabindex -->
        <div 
          bind:this={sidebarItems[index]}
          class="sidebar-item last-item" 
          tabindex="0"
          on:click={() => handleLinkClick(link)}
          on:keydown={(event) => handleKeyDown(event, link)}
        >
          <Icon class="material-icons">{link.icon}</Icon>
          <span class="sidebar-tooltip">{link.text}</span>
        </div>
      {:else}
        <!-- svelte-ignore a11y-no-noninteractive-tabindex -->
        <div 
          bind:this={sidebarItems[index]}
          class="sidebar-item" 
          tabindex="0"
          on:click={() => handleLinkClick(link)}
          on:keydown={(event) => handleKeyDown(event, link)}
        >
          <Icon class="material-icons">{link.icon}</Icon>
          <span class="sidebar-tooltip">{link.text}</span>
        </div>
      {/if}
    {/each}
  </Router>
</div> 

<style>
  .sidebar{
    position: fixed;
    top: 0;
    left: 0;
    height:100vh;
    width: 4rem;
    display: flex;
    flex-direction: column;
    margin: 0;
    background-color: #1e1f22;
    align-items: center;
    z-index: 9999;
  }
  .sidebar-item {
    display: flex;
    position: relative;
    margin-top: .5rem;
    margin-bottom: .5rem;
    align-items: center;
    justify-content: center;
    height: 3rem;
    width: 3rem;
    border-radius: 2rem;
    margin-left: auto;
    margin-right: auto;
    font-size:1.5rem;
    margin-right:.5rem;
    background-color: #313338;
    color: #7289da;
    cursor: pointer;
    box-shadow: 0px 4px 6px -1px rgba(0,0,0,0.1),0px 2px 4px -1px rgba(0,0,0,0.06);
    transition-property: all;
    transition-duration: .15s;
    transition-timing-function: linear;
  }
  .sidebar-item:hover {
    background-color: #7289da;
    color: #FFFFFF;
    border-radius: .75rem;
  }
  .sidebar-tooltip {
    display: none;
    position: absolute; 
    left: 4rem; 
    padding: 0.5rem; 
    margin: 0.5rem; 
    border-radius: 0.375rem; 
    width: auto; 
    min-width: max-content; 
    font-size: 0.75rem;
    line-height: 1rem; 
    font-weight: 700; 
    box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06); 
    padding: 0.5rem;
    background-color: #4B5563;
    color: #FFFFFF;
  }
  .sidebar-item:hover .sidebar-tooltip {
    display: block;
  }
  .divider {
  height: 3px;
  width: 60%;
  margin-left: 0.5rem;
  margin-right: 0.5rem;
  margin-top: auto;
  margin-bottom: .5rem;
  border-radius: 9999px;
  border-width: 1px;
  border-color: #35363C;
  background-color: #35363C;
}
.last-item {
  margin-top: 0.5rem;
}
.active-bar {
  position: absolute;
  left: 0;
  width: 3px;
  height: 3rem;
  background-color: white;
  transition: transform 0.15s ease-in-out;
}
</style>

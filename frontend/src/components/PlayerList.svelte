<script>
    import { onMount } from 'svelte';
    import List, { Item, Graphic, Text, PrimaryText, SecondaryText, Meta } from '@smui/list';
    import { players, removePlayer } from '../store';
  
    let playerList = [];
  
    onMount(() => {
      const unsubscribe = players.subscribe(value => {
        playerList = value;
      });
      return unsubscribe;
    });
  
    function handleRemovePlayer(username) {
      removePlayer(username);
    }
  </script>
  
  <List twoLine avatarList>
    {#each playerList as player}
      <Item>
        <Graphic style={`background-image: url(${player.avatar_url})`} />
        <Text>
          <PrimaryText>{player.username}</PrimaryText>
          <SecondaryText>User ID: {player.userId}</SecondaryText>
        </Text>
        <Meta class="material-icons" on:click={() => handleRemovePlayer(player.username)}>clear</Meta>
      </Item>
    {/each}
  </List>
  
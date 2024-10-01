<script>
    import axios from 'axios';
    import { onMount } from 'svelte';
  
    let fermentations = [];
  
    onMount(async () => {
      const response = await axios.get('http://0.0.0.0:8000/fermentations');
      fermentations = response.data;
    });
  </script>
  
  <table>
    <thead>
      <tr>
        <th>Name</th>
        <th>Start Date</th>
        <th>Bottled At Date</th>
        <th>Recipe Notes</th>
        <th>Tasting Notes</th>
      </tr>
    </thead>
    <tbody>
      {#each fermentations as f}
        <tr>
          <td>{f.id}</td>
          <td>{f.name}</td>
          <td>{f.start_date}</td>
          <td>{f.bottled_at}</td>
          <td>{f.recipe_notes}</td>
          <td>{f.tasting_notes}</td>
        </tr>
      {/each}
    </tbody>
  </table>
  
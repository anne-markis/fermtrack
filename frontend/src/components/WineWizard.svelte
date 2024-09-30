<script>
    import axios from 'axios';
    let question = '';
    let advice = 'Ask the wine wizard anything';
    let thinking = false;
  
    const submitQuestion = async () => {
      thinking = true;
      advice = '';
      try {
        const response = await axios.post('http://localhost:8000/fermentations/advice', { question });
        advice = response.data.advice;
      } catch (error) {
        advice = 'An error occurred.';
      } finally {
        thinking = false;
      }
    };
  </script>
  
  <style>
    .container {
      display: flex;
    }
    .left, .right {
      flex: 1;
      padding: 20px;
    }
    .right {
      border-left: 1px solid #ccc;
    }
  </style>
  
  <div class="container">
    <div class="left">
      <textarea bind:value={question} maxlength="500" rows="10" cols="30"></textarea>
      <br/>
      <button on:click={submitQuestion}>Submit</button>
    </div>
    <div class="right">
      {#if thinking}
        <p>Thinking...</p>
      {:else}
        <p>{advice}</p>
      {/if}
    </div>
  </div>
  
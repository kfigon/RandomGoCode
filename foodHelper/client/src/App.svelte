<script>
	let ingredient = ""
	let items = [];
	let suggestedIngredients = [];

	function inputHandler(e) {
		ingredient = e.target.value;

		fetch('http://localhost:8000/api/ingredients?prefix='+ingredient)
				.then(response => response.json())
				.then(response => response.ingredients)
				.then(data => {
					suggestedIngredients = data;
					console.log(suggestedIngredients)
				});
	}

	
</script>

<main>
	<h1>Hello! What do you have in your fridge?</h1>

	<input type="text" value={ingredient} on:input={inputHandler}>
	
	{#if suggestedIngredients.length !== 0}
	<p>?!</p>
	<ul>
		{#each suggestedIngredients as item}
			<li>{item}</li>
		{/each}
	</ul>
	{/if}

	<ul>
		{#each items as item}
			<li>{item}</li>
		{/each}
	</ul>
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 3em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>
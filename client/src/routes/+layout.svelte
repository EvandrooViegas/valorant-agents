<script>
	import Navbar from "$components/navbar/Navbar.svelte";
  import "../app.css";
</script>

<main>
  <Navbar />
  <div class="p-5">
    <slot />
  </div>
</main>
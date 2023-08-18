<script>
	import Navbar from "$components/navbar/Navbar.svelte";
	import Toaster from "$components/toaster/Toaster.svelte";
  import "../app.css";
</script>

<main>
  <Navbar />
  <Toaster />
  <div class="p-5">
    <slot />
  </div>
</main>
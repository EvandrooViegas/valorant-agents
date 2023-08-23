<script>
	import Navbar from "../lib/components/navbar/Navbar.svelte";
	import Toaster from "../lib/components/toaster/Toaster.svelte";
  import "../app.css";
</script>

<main>
  <Navbar />
  <Toaster />
  <div class="p-5">
    <slot />
  </div>
</main>
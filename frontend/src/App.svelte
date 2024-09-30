<script>
	import Header from './components/Header.svelte';
	import Navbar from './components/Navbar.svelte';
	import Profile from './components/Profile.svelte';
	import Fermentations from './components/Fermentations.svelte';
	import WineWizard from './components/WineWizard.svelte';
	import { Router, Route, Link } from 'svelte-routing';
  </script>
  
  <Header />
  <Navbar />
  
  <Router>
	<Route path="/" component={Profile} />
	<Route path="/fermentations" component={Fermentations} />
	<Route path="/wine-wizard" component={WineWizard} />
  </Router>
  
<script lang="ts">
  import { Router, Route, navigate } from "svelte-navigator";
  import SignUp from "./lib/SignUp.svelte";
  import "./index.css"
import LoginIn from "./lib/LoginIn.svelte";
import Home from "./lib/Home.svelte"

  let token = localStorage.getItem("token")
</script>

<Router>
  <Route path="/">
    <Home />
    {#if token==null}
    {navigate("/signup", {replace: true})}
    {/if}
  </Route>
  <Route path="/signup">
    <SignUp />
  </Route>
  <Route path="/login">
    <LoginIn />
  </Route>
</Router>

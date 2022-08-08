<script lang="ts">
  import axios from 'axios'
  import { Router, Route, navigate } from "svelte-navigator";
  import SignUp from "./lib/SignUp.svelte";
  import "./index.css"
import LoginIn from "./lib/LoginIn.svelte";
import Home from "./lib/Home.svelte"
import {SERVER_URL} from "../../secrets"

axios.defaults.baseURL = SERVER_URL

</script>

<Router>
  <Route path="/">
    <Home />
    {#if localStorage.getItem("token")==null}
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

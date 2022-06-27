<script lang="ts">
import {navigate} from 'svelte-navigator'
import {onMount} from 'svelte'
import axios from 'axios';
import AddSecret from './AddSecret.svelte';
import Otps from './Otps.svelte';
import type {Response} from "../Types"
import {AllData} from "../stores"

if(localStorage.getItem("token") == null) {
    navigate("/", {replace: true})
}

let token = localStorage.getItem("token")
let data: Partial<Response> = {"issuers": {}}
let ifError = false
let errorString = ""
let loading = true

onMount(() => {
    axios.get("/codes", {
        headers: {
            "Authorization": ` Bearer ${token}`
        }
    })
    .then(d => {
        data = d.data;
        loading=false;
        AllData.update(() => d.data)
    })
    .catch(err => {
        if(err.response) {
            if (err.response.status == 400) {
            console.log(err)
            if(err.response.data=="User has no secrets yet!") {
                errorString = err.response.data
                ifError == true
            } else if (err.response.data=="User doesn't exist") {
                localStorage.removeItem("token")
                navigate("/login", {replace:true})
            }
        }
        } else {
            console.log(err)
        }
        
    })
})

AllData.subscribe(v => data=v)

const logout = () => {
    localStorage.removeItem("token")
    navigate("/login", {replace:true})
}
</script>

<main class="flex flex-col items-center justify-center">
    <!-- <h1 class="text-5xl font-bold pb-6">Hello There</h1> -->
    <div class="dropdown">
        <!-- svelte-ignore a11y-label-has-associated-control -->
        <label tabindex="0" class="btn text-5xl font-bold pb-6 m-1">Hello There</label>
        <ul tabindex="0" class="dropdown-content menu p-2 shadow bg-base-100 rounded-box w-52">
          <li><button class="btn btn-sm btn-secondary" on:click={logout}>Logout</button></li>
        </ul>
      </div>
    <div class="flex flex-wrap">
        {#if ifError}
            {errorString}
        {/if}
    {#if !loading && !ifError}
        {#each Object.entries(data.issuers) as [issuerN, accounts] (issuerN)}
            <h2 class="text-xl text-bold">{issuerN}:</h2><br/>
            <Otps accountData={accounts} />
        {/each}
    {/if}
</div>
<div class="py-6">
    <AddSecret/>
</div>
</main>
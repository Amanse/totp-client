<script lang="ts">
import {navigate} from 'svelte-navigator'
import {onMount} from 'svelte'
import axios from 'axios';
import AddSecret from './AddSecret.svelte';
import Otps from './Otps.svelte';

if(localStorage.getItem("token") == null) {
    navigate("/", {replace: true})
}

let token = localStorage.getItem("token")
let data = {}
let loading = true

onMount(() => {
    axios.get("http://localhost:8080/codes", {
        headers: {
            "Authorization": ` Bearer ${token}`
        }
    })
    .then(d => {data = d.data ;loading=false})
    .catch(err => {
        if(err.response) {
            if (err.response.status == 400) {
            console.log(err)
            if(err.response.data=="User has no secrets yet!") {
                data="User has no secrets yet!"
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
</script>

<main class="flex flex-col items-center justify-center">
    <h1 class="text-5xl font-bold pb-6">Hello There</h1>
    <div class="flex flex-wrap">
    {#if !loading}
        {#each Object.entries(data.issuers) as [issuerN, accounts] (issuerN)}
            <h2 class="text-xl text-bold">{issuerN}:</h2><br/>
            <Otps accountData={accounts} />
        {/each}
    {/if}
</div>
<div class="py-6">
    <AddSecret  fullData = {data}/>
</div>
</main>
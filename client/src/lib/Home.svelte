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
let time: number = 30 - Math.round(new Date().valueOf() / 1000) % 30

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
        } else if(err.response.status == 401) {
            localStorage.removeItem("token")
            navigate("/login", {replace: true})
        }
        } else {
            console.log(err)
        }
        
    })
})

setInterval(() => time = (30 - Math.round(new Date().valueOf() / 1000) % 30)
, 1000)


AllData.subscribe(v => data=v)

const logout = () => {
    localStorage.removeItem("token")
    navigate("/login", {replace:true})
}

const handleDelete = () => {
    let accountToD = sessionStorage.getItem("accountToD")
    let secretToD = sessionStorage.getItem("secretToD")
    let issuerToD = sessionStorage.getItem("issuerToD")

    console.log(`Delete ${accountToD} with ${secretToD} and ${issuerToD}`)
    let accounts = data["issuers"][issuerToD]["accounts"]
    delete accounts[accountToD]
    let issuers = data["issuers"]

    if (Object.keys(accounts).length === 0) {
        console.log("should delete issuer")
        delete issuers[issuerToD]
        data["issuers"] = issuers
        
    } else {
        data["issuers"][issuerToD]["accounts"] = accounts
    }
    console.log(data)

    axios.post("codes", {issuers: data.issuers}, {
        headers: {
            'Authorization': ` Bearer ${token}`
        }
    })

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
    <div class="grid grid-cols-2 md:grid-cols-4 gap-7">
        {#if ifError}
            {errorString}
        {/if}
    {#if !loading && !ifError}
        {#each Object.entries(data.issuers) as [issuerN, accounts] (issuerN)}
        <div>
            <Otps accountData={accounts} {issuerN} {handleDelete} />
        </div>
        {/each}
    {/if}
</div>
<!-- <span class="countdown">
    <span style="--value:{time};"></span>
  </span> -->
  <progress class="progress progress-secondary w-56" value={time} max="30"></progress>

<div class="py-6">
    <AddSecret/>
</div>
</main>
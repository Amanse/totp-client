<script lang="ts">
import axios from "axios";
import {onMount} from "svelte"
import {secretData} from "../stores"

interface Response {
    issuers: Issuer
}

interface Issuer {
    [name: string]: Accounts
}

interface Accounts {
    accounts: Account
}

interface Account {
    [name: string]: string
}

let issuerN: string = ""
let accountName: string = ""
let secret: string = ""

let token = localStorage.getItem("token")

let makeNew = false

export let fullData
console.log(fullData)
const handleAddSec = () => {
    if(fullData == {}){
            fullData["issuers"][issuerN] = {"accounts": {}}
            fullData["issuers"][issuerN]["accounts"][accountName] = secret
        }else {
            if(fullData["issuers"][issuerN] != null) {
                fullData["issuers"][issuerN]["accounts"][accountName] = secret
            } else {
                fullData["issuers"][issuerN] = {"accounts": {}}
                fullData["issuers"][issuerN]["accounts"][accountName] = secret
        }
    }
   
    let fis = fullData.issuers
    console.log(fullData.issuers)
    axios.post("http://localhost:8080/codes", {issuers: fis}, {
            headers: {
                'Authorization': ` Bearer ${token}`
            }
        }
    ).then(data => console.log(`data is ${data}`))
    .catch(err => console.log(err))
    console.log("Add")
}

</script>

<div class="card w-96 bg-primary text-primary-content">
    <div class="card-body">
      <h2 class="card-title">Add secret</h2>
      <input type="text" placeholder="Issuer" bind:value={issuerN} class="input input-bordered input-secondary w-full max-w-xs" />
      <input type="text" placeholder="Account" bind:value={accountName} class="input input-bordered input-secondary w-full max-w-xs" />
      <input type="text" placeholder="Secret" bind:value={secret} class="input input-bordered input-secondary w-full max-w-xs" />
      <div class="card-actions justify-end">
        <button class="btn btn-secondary" on:click={handleAddSec}>Add secert</button>
      </div>
    </div>
  </div>
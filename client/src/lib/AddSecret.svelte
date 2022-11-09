<script lang="ts">
import axios from "axios";
import type {Response} from "../Types"
import {AllData} from "../stores"
import {onMount} from "svelte"

let issuerN: string = ""
let accountName: string = ""
let secret: string = ""

let token = localStorage.getItem("token")

let fullData: Partial<Response>
console.log(fullData)
const handleAddSec = () => {
    if(fullData["issuers"][issuerN] == null) {
        fullData["issuers"][issuerN] = {"accounts": {}}
        fullData["issuers"][issuerN]["accounts"][accountName] = secret.replace(/ /g,'')
    } else {
        fullData["issuers"][issuerN]["accounts"][accountName] = secret.replace(/ /g,'')
    }
   
    let fis = fullData.issuers
    console.log(fullData.issuers)
    axios.post("codes", {issuers: fis}, {
            headers: {
                'Authorization': ` Bearer ${token}`
            }
        }
    ).then(() => {
        AllData.update(() => fullData)
        issuerN = ""
        accountName = ""
        secret = ""
    })
    .catch(err => console.log(err))
    console.log("Add")
}

const updateFullData = () => {
    AllData.subscribe(v => {
        fullData = v
    })
}

onMount(() => {
   updateFullData()
})

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

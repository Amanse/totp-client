<script lang="ts">
import totp from "totp-generator";
import type {Accounts} from "../Types"

let time: number;

let secretToD: string;
let accountToD: string;

setInterval(() =>handleTime(), 1000)

const handleTime = () => {
    time = (30 - Math.round(new Date().valueOf() / 1000) % 30)
    if( time==30 ){
        accountData = accountData
    }
}

const copySecretToClip = (secret) => {
    let text = totp(secret)
    navigator.clipboard.writeText(text).then(() => {}, (err) => console.log(err))
}

const propsToSendUp = (accountN, secret, issu) => {
  console.log("called with: ", accountN)
  sessionStorage.setItem("secretToD", secret)
  sessionStorage.setItem("accountToD", accountN)
  sessionStorage.setItem("issuerToD", issu)
}

export let accountData: Accounts;
export let handleDelete;
export let issuerN;
</script>

<main>
  <h2 class="text-xl text-bold">{issuerN}:</h2>

    {#each Object.entries(accountData.accounts) as [accountN, secret](accountN)}
        <h5 class="text-xl text-secondary">{accountN}</h5>
        <div class="flex flex-row">
        <p class="py-2 text-accent" on:click={() => copySecretToClip(secret)}>Secret: {totp(secret)}</p>
        
        <div on:click={() => propsToSendUp(accountN, secret, issuerN)}>
        <label for="my-modal" class="modal-button">
            <svg  xmlns="http://www.w3.org/2000/svg" viewBox="0 0 48 48" width="32px" height="32px"><path fill="#b39ddb" d="M30.6,44H17.4c-2,0-3.7-1.4-4-3.4L9,11h30l-4.5,29.6C34.2,42.6,32.5,44,30.6,44z"/><path fill="#9575cd" d="M28 6L20 6 14 12 34 12z"/><path fill="#7e57c2" d="M10,8h28c1.1,0,2,0.9,2,2v2H8v-2C8,8.9,8.9,8,10,8z"/></svg>
        </label>
        </div>
    
    </div>
    {/each}
    <!-- The button to open modal -->

<!-- Put this part before </body> tag -->
<input type="checkbox" id="my-modal" class="modal-toggle" />
<div class="modal">
  <div class="modal-box">
    <h3 class="font-bold text-lg">Are you sure you want to delete?</h3>
    <div class="modal-action">
      <label for="my-modal" on:click={() => handleDelete()} class="btn btn-warning">Yes</label>
      <label for="my-modal" class="btn btn-info">No</label>
    </div>
  </div>
</div>
</main>
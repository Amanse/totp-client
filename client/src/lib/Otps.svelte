<script lang="ts">
import totp from "totp-generator";
import type {Accounts} from "../Types"

let time: number;

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

export let accountData: Accounts;
</script>

<main>
    {#each Object.entries(accountData.accounts) as [accountN, secret](accountN)}
        <h5 class="text-xl text-secondary">{accountN}</h5>
        <div class="flex flex-row">
        <p class="py-2 text-accent" on:click={() => copySecretToClip(secret)}>Secret: {totp(secret)}</p>
        </div>
    {/each}
    
</main>
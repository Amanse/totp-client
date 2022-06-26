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

export let accountData: Accounts;
</script>

<main>
    {#each Object.entries(accountData.accounts) as [accountN, secret](accountN)}
        <h5 class="text-xl px-4 text-secondary">{accountN}</h5>
        <p class="py-2 text-accent">Secret: {totp(secret)}</p>
        <span class="countdown">
            <span style="--value:{time};"></span>
          </span>
    {/each}
</main>
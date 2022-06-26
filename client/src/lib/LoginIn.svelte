<script lang="ts">
    import axios from 'axios';
    import {Link} from 'svelte-navigator'
    let username = "";
    let password = "";

    let errorString = "";
    let ifError = false;
    let ifSuccess = false;

    const handleSignUp = async () => {
        console.log("called")
        if (username == "" || password == "") {
          errorString = "Please input all fields!"
          ifError = true;
          return
        }

        username = username.trim()
        password = password.trim()

        try {
            axios.post("http://localhost:8080/login", {
            username,
            password
        }).then(data => {
            localStorage.setItem("token", data.data.token)
            ifError = false
            errorString = ""
            ifSuccess = true
        }).catch(error => {
            console.log(error.response.status)
            if(error.response.status == 400) {
                errorString = error.response.data
                ifError = true
            }
        })
        } catch(error) {
            console.log(error)
        }   
    }
</script>

<main>
    <!-- <input type="text" bind:value={username} placeholder="username" />
    <input type="text" bind:value={password} placeholder="username" />

    <button on:click={handleSignUp} >Sign up</button> -->

    <div class="hero min-h-screen bg-base-200">
        <div class="hero-content flex-col lg:flex-row-reverse">
          <div class="text-center lg:text-left">
            <h1 class="text-5xl font-bold">Login</h1>
            <p class="py-6">Login Now to get back your TOTPs from the most trusted platform!</p>
          </div>
          <div class="card flex-shrink-0 w-full max-w-sm shadow-2xl bg-base-100">
            <div class="card-body">
              {#if ifSuccess}
                <!-- svelte-ignore a11y-label-has-associated-control -->
                <label class="label">
                    <p class="label-text text-green-500">Success! Welcome back</p>
                </label>
                {/if}
                {#if ifError}
                <!-- svelte-ignore a11y-label-has-associated-control -->
                <label class="label">
                    <p class="label-text text-red-500">{errorString}</p>
                </label>
                {/if}
              <div class="form-control">
                <input type="text" placeholder="Username" bind:value={username} class="input py-3 input-secondary input-bordered" />
              </div>
              <br />
              <div class="form-control">
                <input type="password" placeholder="Password" bind:value={password} class="input input-secondary input-bordered" />
              </div>
              <div class="form-control mt-6">
                <button class="btn btn-secondary" on:click={handleSignUp}>Login</button>
              </div>
              <!-- svelte-ignore a11y-label-has-associated-control -->
              <label class="label">
                <Link to="/signup" class="label-text-alt link link-hover">Don't have an account? Signup</Link>
              </label>
            </div>
          </div>
        </div>
      </div>
</main>
<script lang="ts">
    import axios from 'axios';
    import {Link, navigate} from 'svelte-navigator'
    let username = "";
    let password = "";


    let ifError = false;
    let errorString = "";
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
            axios.post("/signup", {
            username,
            password
        }).then(data => {
            localStorage.setItem("token", data.data.token)
            ifError = false
            errorString = ""
            ifSuccess = true
            navigate("/login", {replace:true})
        }).catch(err => {
          if (err.response.status == 400) {
            errorString = err.response.data
            ifError = true;
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
            <h1 class="text-5xl font-bold">Signup</h1>
            <p class="py-6">Signup today to start syning your TOTPs accross devices with full security!</p>
          </div>
          <div class="card flex-shrink-0 w-full max-w-sm shadow-2xl bg-base-100">
            <div class="card-body">
              {#if ifSuccess}
              <!-- svelte-ignore a11y-label-has-associated-control -->
              <label class="label">
                  <p class="label-text text-green-500">Success! Redirecting to login</p>
              </label>
              {/if}
              {#if ifError}
              <!-- svelte-ignore a11y-label-has-associated-control -->
              <label class="label">
                  <p class="label-text text-red-500">{errorString}</p>
              </label>
              {/if}
              <div class="form-control">
                <input type="text" placeholder="Username" bind:value={username} class="input input-secondary input-bordered" />
              </div>
              <br />
              <div class="form-control">
                <input type="password" placeholder="Password" bind:value={password} class="input input-secondary input-bordered" />
              </div>
              <div class="form-control mt-6">
                <button class="btn btn-secondary" on:click={handleSignUp}>Signup</button>
              </div>
              <!-- svelte-ignore a11y-label-has-associated-control -->
              <label class="label">
                <Link to="/login" class="label-text-alt link link-hover">Already have an account? Login</Link>
              </label>
            </div>
          </div>
        </div>
      </div>
</main>
import { component$, useSignal, $ } from "@builder.io/qwik";
import { useNavigate, Link } from "@builder.io/qwik-city";
import { Client } from "@passwordlessdev/passwordless-client";
import axios from "axios";
import { sha256 } from "js-sha256";

export default component$(() => {
  let username = useSignal("");
  let password = useSignal("");
  let errorString = useSignal("");
  let ifError = useSignal(false);

  const nav = useNavigate();

  const loginPasskey = $(() => {
    const p = new Client({
      apiKey: "totp:public:4f46665f2e6a41f3abdbdae8ef539114",
    });

    p.signinWithId(username.value).then(({ token, error }) => {
      if (error) {
        console.log(error);
        return;
      }

      axios
        .post("http://localhost:8080/login/passkey", { token })
        .then((data) => {
          localStorage.setItem("token", data.data.token);
          ifError.value = false;
          errorString.value = "";
          nav("/");
        });
    });
  });

  const login = $(() => {
    try {
      axios
        .post("http://localhost:8080/login", {
          username: username.value.trim(),
          password: sha256(password.value.trim()),
        })
        .then((data) => {
          localStorage.setItem("token", data.data.token);
          ifError.value = false;
          errorString.value = "";
          nav("/");
        })
        .catch((err) => {
          if (err.response.status == 400) {
            errorString.value = err.response.data;
            ifError.value = true;
          }
        });
    } catch (error) {
      console.log(error);
    }
  });
  return (
    <>
      <div class="hero min-h-screen bg-base-200">
        <div class="hero-content flex-col lg:flex-row-reverse">
          <div class="text-center lg:text-left">
            <h1 class="text-5xl font-bold">Login</h1>
            <p class="py-6">Welcome back!</p>
          </div>
          <div class="card flex-shrink-0 w-full max-w-sm shadow-2xl bg-base-100">
            <div class="card-body">
              {ifError && (
                <label class="label">
                  <p class="label-text text-red-500">{errorString}</p>
                </label>
              )}
              <div class="form-control">
                <input
                  type="text"
                  placeholder="Username"
                  bind:value={username}
                  class="input input-secondary input-bordered"
                />
              </div>
              <br />
              <div class="form-control">
                <input
                  type="password"
                  placeholder="Password"
                  bind:value={password}
                  class="input input-secondary input-bordered"
                />
              </div>
              <div class="form-control mt-6">
                <button class="btn btn-secondary" onClick$={() => login()}>
                  Login
                </button>
                <button
                  class="btn btn-secondary mt-2"
                  onClick$={() => loginPasskey()}
                >
                  Login with passkey
                </button>
              </div>
              <label class="label">
                <Link href="/signup" class="label-text-alt link link-hover">
                    Don't have an account? Sign up!
                </Link>
              </label>
            </div>
          </div>
        </div>
      </div>
    </>
  );
});

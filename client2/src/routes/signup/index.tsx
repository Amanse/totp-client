import { component$, useSignal, $ } from "@builder.io/qwik";
import { useNavigate, Link } from "@builder.io/qwik-city";
import axios from "axios";
import { Client } from "@passwordlessdev/passwordless-client";
import { sha256 } from "js-sha256";

export default component$(() => {
  let username = useSignal("");
  let password = useSignal("");
  let errorString = useSignal("");
  let ifError = useSignal(false);

  const nav = useNavigate();

  const signUpPasskey = $(async () => {
    const p = new Client({
      apiKey: "totp:public:4f46665f2e6a41f3abdbdae8ef539114",
    });

    axios
      .post("http://localhost:8080/signup/passkey", {
        username: username.value,
        password: "hehe",
      })
      .then((r) => {
        p.register(r.data.token).then(({ token, error }) => {
          if (error) {
            console.log(error);
          }
          if (token) {
            localStorage.setItem("token", token);
            nav("/login");
          }
        });
      });
  });

  const signUp = $(() => {
    username.value = username.value.trim();
    password.value = sha256(password.value.trim());
    try {
      axios
        .post("http://localhost:8080/signup", {
          username: username.value,
          password: password.value,
        })
        .then((data) => {
          localStorage.setItem("token", data.data.token);
          ifError.value = false;
          errorString.value = "";
          nav("/login");
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
            <h1 class="text-5xl font-bold">Signup</h1>
            <p class="py-6">
              Signup today to start syning your TOTPs accross devices with full
              security!
            </p>
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
                <button class="btn btn-secondary" onClick$={() => signUp()}>
                  Signup
                </button>
                <button
                  class="btn btn-secondary mt-2"
                  onClick$={() => signUpPasskey()}
                >
                  Signup with passkey
                </button>
              </div>
              <label class="label">
                <Link href="/login" class="label-text-alt link link-hover">
                  Already have an account? Login
                </Link>
              </label>
            </div>
          </div>
        </div>
      </div>
    </>
  );
});

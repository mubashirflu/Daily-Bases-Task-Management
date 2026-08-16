<!-- <script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
const email = ref("");
const password = ref("");
const router=useRouter()
async function login() {
  const response = await fetch("http://localhost:8080/api/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      email: email.value,
      password: password.value,
    }),
  });

  const data = await response.json();

    if (response.ok) {
    router.push("/dashboard");
  } else {
    alert(data.error || "Login failed");
  }
}
</script>

<template>
  <div>
    <h1>Login</h1>

    <input
      v-model="email"
      type="email"
      placeholder="Email"
    />

    <input
      v-model="password"
      type="password"
      placeholder="Password"
    />

    <button @click="login">
      Login
    </button>
  </div>
</template> -->
<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";

const router = useRouter();

const email = ref("");
const password = ref("");
const loading = ref(false);
const errorMsg = ref("");

async function login() {
  errorMsg.value = "";
  loading.value = true;

  try {
    const response = await fetch("http://localhost:8080/api/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        email: email.value,
        password: password.value,
      }),
    });

    const data = await response.json();

    if (!response.ok) {
      errorMsg.value = data.error || "Couldn't log you in. Check your details and try again.";
      return;
    } 
     localStorage.setItem("token", response.data.token);

    // console.log("TOKEN SAVED:", localStorage.getItem("token"));
     router.push("/dashboard");
  } catch (error) {
    errorMsg.value = "Couldn't reach the server. Check your connection and try again.";
  } finally {
    loading.value = false;
  }
}
</script>

<template>
  <div class="auth-page">
    <!-- Brand panel -->
    <section class="auth-brand">
      <div class="auth-brand-inner">
        <div class="brand-mark">
          <span class="brand-dot"></span>
          Tally
        </div>

        <h1 class="brand-headline">Small steps,<br />counted.</h1>
        <p class="brand-sub">
          Tally keeps your day honest — one task, one check, one step forward.
        </p>

        <svg class="brand-art" viewBox="0 0 320 220" fill="none" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
          <rect x="14" y="24" width="292" height="56" rx="14" fill="rgba(255,255,255,0.10)" />
          <circle cx="40" cy="52" r="12" fill="rgba(255,255,255,0.85)" />
          <path d="M35 52L39 56L47 47" stroke="#5B5BD6" stroke-width="3" stroke-linecap="round" stroke-linejoin="round"/>
          <rect x="64" y="46" width="150" height="6" rx="3" fill="rgba(255,255,255,0.55)" />
          <rect x="64" y="58" width="90" height="6" rx="3" fill="rgba(255,255,255,0.30)" />

          <rect x="14" y="92" width="292" height="56" rx="14" fill="rgba(255,255,255,0.10)" />
          <circle cx="40" cy="120" r="12" fill="rgba(255,255,255,0.30)" stroke="rgba(255,255,255,0.6)" stroke-width="2"/>
          <rect x="64" y="114" width="170" height="6" rx="3" fill="rgba(255,255,255,0.55)" />
          <rect x="64" y="126" width="110" height="6" rx="3" fill="rgba(255,255,255,0.30)" />

          <rect x="14" y="160" width="292" height="56" rx="14" fill="rgba(255,255,255,0.10)" />
          <circle cx="40" cy="188" r="12" fill="rgba(255,255,255,0.30)" stroke="rgba(255,255,255,0.6)" stroke-width="2"/>
          <rect x="64" y="182" width="130" height="6" rx="3" fill="rgba(255,255,255,0.55)" />
          <rect x="64" y="194" width="80" height="6" rx="3" fill="rgba(255,255,255,0.30)" />
        </svg>
      </div>
    </section>

    <!-- Form panel -->
    <section class="auth-form-panel">
      <div class="auth-form-card">
        <h2>Welcome back</h2>
        <p class="auth-form-sub">Log in to pick up where you left off.</p>

        <p v-if="errorMsg" class="auth-error" role="alert">{{ errorMsg }}</p>

        <form @submit.prevent="login" novalidate>
          <div class="field">
            <label for="email">Email</label>
            <input
              id="email"
              v-model="email"
              type="email"
              autocomplete="email"
              placeholder="you@example.com"
              required
            />
          </div>

          <div class="field">
            <label for="password">Password</label>
            <input
              id="password"
              v-model="password"
              type="password"
              autocomplete="current-password"
              placeholder="••••••••"
              required
            />
          </div>

          <button type="submit" class="btn-primary" :disabled="loading">
            <span v-if="loading" class="spinner" aria-hidden="true"></span>
            {{ loading ? "Logging in…" : "Log in" }}
          </button>
        </form>

        <p class="auth-switch">
          New to Tally?
          <router-link to="/register">Create an account</router-link>
        </p>
      </div>
    </section>
  </div>
</template>

<style scoped>
.auth-page {
  min-height: 100vh;
  display: grid;
  grid-template-columns: 1fr 1fr;
}

/* Brand panel */
.auth-brand {
  background: linear-gradient(160deg, #4949BE 0%, #5B5BD6 45%, #7A6FE0 100%);
  color: #fff;
  display: flex;
  align-items: center;
  padding: 56px;
}

.auth-brand-inner {
  max-width: 420px;
}

.brand-mark {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  font-family: var(--font-display);
  font-weight: 700;
  font-size: 20px;
  letter-spacing: -0.01em;
  margin-bottom: 48px;
}

.brand-dot {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  background: #fff;
  box-shadow: 0 0 0 4px rgba(255, 255, 255, 0.25);
}

.brand-headline {
  font-size: 40px;
  line-height: 1.15;
  font-weight: 600;
  letter-spacing: -0.02em;
  margin-bottom: 16px;
}

.brand-sub {
  font-size: 16px;
  line-height: 1.6;
  color: rgba(255, 255, 255, 0.82);
  margin: 0 0 40px;
  max-width: 340px;
}

.brand-art {
  width: 100%;
  height: auto;
}

/* Form panel */
.auth-form-panel {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 32px;
}

.auth-form-card {
  width: 100%;
  max-width: 380px;
}

.auth-form-card h2 {
  font-size: 26px;
  font-weight: 600;
  letter-spacing: -0.01em;
}

.auth-form-sub {
  color: var(--ink-soft);
  font-size: 15px;
  margin: 8px 0 28px;
}

.auth-error {
  background: var(--danger-soft);
  color: var(--danger);
  font-size: 14px;
  padding: 12px 14px;
  border-radius: var(--radius-sm);
  margin: 0 0 20px;
}

.field {
  margin-bottom: 18px;
}

.field label {
  display: block;
  font-size: 13px;
  font-weight: 600;
  color: var(--ink-soft);
  margin-bottom: 7px;
}

.field input {
  width: 100%;
  padding: 12px 14px;
  font-size: 15px;
  font-family: var(--font-body);
  color: var(--ink);
  background: var(--surface);
  border: 1.5px solid var(--border);
  border-radius: var(--radius-sm);
  transition: border-color 0.15s ease, box-shadow 0.15s ease;
}

.field input::placeholder {
  color: var(--ink-faint);
}

.field input:focus {
  outline: none;
  border-color: var(--accent);
  box-shadow: 0 0 0 4px var(--accent-soft);
}

.btn-primary {
  width: 100%;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  margin-top: 6px;
  padding: 13px 18px;
  font-size: 15px;
  font-weight: 600;
  color: #fff;
  background: var(--accent);
  border: none;
  border-radius: var(--radius-sm);
  cursor: pointer;
  transition: background 0.15s ease, transform 0.1s ease;
}

.btn-primary:hover:not(:disabled) {
  background: var(--accent-hover);
}

.btn-primary:active:not(:disabled) {
  transform: scale(0.98);
}

.btn-primary:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.spinner {
  width: 15px;
  height: 15px;
  border: 2px solid rgba(255, 255, 255, 0.4);
  border-top-color: #fff;
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

.auth-switch {
  margin-top: 26px;
  font-size: 14px;
  color: var(--ink-soft);
  text-align: center;
}

.auth-switch a {
  color: var(--accent);
  font-weight: 600;
  text-decoration: none;
}

.auth-switch a:hover {
  text-decoration: underline;
}

/* Mobile */
@media (max-width: 860px) {
  .auth-page {
    grid-template-columns: 1fr;
  }

  .auth-brand {
    padding: 40px 28px 32px;
  }

  .auth-brand-inner {
    max-width: none;
  }

  .brand-mark {
    margin-bottom: 24px;
  }

  .brand-headline {
    font-size: 30px;
  }

  .brand-sub {
    margin-bottom: 24px;
  }

  .brand-art {
    display: none;
  }

  .auth-form-panel {
    padding: 32px 24px 56px;
  }
}
</style>
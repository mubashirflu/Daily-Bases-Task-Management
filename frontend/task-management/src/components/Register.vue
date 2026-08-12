<!-- 
<script setup>
import { ref } from 'vue'
import { registerUser } from '@/services/api'
import { useRouter } from 'vue-router'

const router = useRouter()

const name = ref('')
const email = ref('')
const password = ref('')

async function register() {
  try {
    const data = await registerUser({
      name: name.value,
      email: email.value,
      password: password.value,
    })
    router.push('/login')
  } catch (error) {
    console.error(error)
    alert(error.message)
  }
}
</script>

<template>
  <div>
    <h1>USER REGISTRATION SCREEN!!!</h1>

    <input
      v-model="name"
      placeholder="Enter your Name:"
    />

    <input
      v-model="email"
      placeholder="Enter your Email:"
    />

    <input
      v-model="password"
      type="password"
      placeholder="Enter Password:"
    />

    <button @click="register">
      Register
    </button>
  </div>
</template>
``` -->

<script setup>
import { ref } from "vue";
import { registerUser } from "@/services/api";
import { useRouter } from "vue-router";

const router = useRouter();

const name = ref("");
const email = ref("");
const password = ref("");
const loading = ref(false);
const errorMsg = ref("");

async function register() {
  errorMsg.value = "";
  loading.value = true;

  try {
    await registerUser({
      name: name.value,
      email: email.value,
      password: password.value,
    });

    router.push("/login");
  } catch (error) {
    console.error(error);
    errorMsg.value = error.message || "Couldn't create your account. Please try again.";
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

        <h1 class="brand-headline">Start your<br />first tally.</h1>
        <p class="brand-sub">
          A clear list, a clean count, a day you can actually finish.
        </p>

        <svg class="brand-art" viewBox="0 0 320 220" fill="none" xmlns="http://www.w3.org/2000/svg" aria-hidden="true">
          <circle cx="160" cy="110" r="72" fill="rgba(255,255,255,0.08)" />
          <circle cx="160" cy="110" r="72" stroke="rgba(255,255,255,0.35)" stroke-width="2" stroke-dasharray="6 8" />
          <path
            d="M118 112 L148 142 L206 82"
            stroke="#fff"
            stroke-width="10"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
      </div>
    </section>

    <!-- Form panel -->
    <section class="auth-form-panel">
      <div class="auth-form-card">
        <h2>Create your account</h2>
        <p class="auth-form-sub">Takes less than a minute — no credit card, no checklist.</p>

        <p v-if="errorMsg" class="auth-error" role="alert">{{ errorMsg }}</p>

        <form @submit.prevent="register" novalidate>
          <div class="field">
            <label for="name">Name</label>
            <input
              id="name"
              v-model="name"
              type="text"
              autocomplete="name"
              placeholder="Jane Cooper"
              required
            />
          </div>

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
              autocomplete="new-password"
              placeholder="At least 8 characters"
              required
            />
          </div>

          <button type="submit" class="btn-primary" :disabled="loading">
            <span v-if="loading" class="spinner" aria-hidden="true"></span>
            {{ loading ? "Creating account…" : "Create account" }}
          </button>
        </form>

        <p class="auth-switch">
          Already have an account?
          <router-link to="/login">Log in</router-link>
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
  max-width: 280px;
  height: auto;
  display: block;
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
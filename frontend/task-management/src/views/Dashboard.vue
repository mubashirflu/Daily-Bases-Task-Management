<script setup>
import { onMounted, ref, computed } from "vue";
import { useRouter } from "vue-router";

const tasks = ref([]);
const loadingTasks = ref(true);

const title = ref("");
const description = ref("");
const adding = ref(false);
const errorMsg = ref("");

const total = computed(() => tasks.value.length);

const completed = computed(() =>
  tasks.value.filter((task) => task.complete).length
);

const percent = computed(() => {
  if (total.value === 0) {
    return 0;
  }

  return Math.round((completed.value / total.value) * 100);
});

// =========================
// SVG RING
// =========================

const ringRadius = 22;
const ringCircumference = 2 * Math.PI * ringRadius;

const ringOffset = computed(() => {
  return ringCircumference -
    (percent.value / 100) * ringCircumference;
});

// =========================
// TOKEN
// =========================

function getToken() {
  return localStorage.getItem("token");
}

// =========================
// GET TASKS
// =========================

async function getTasks() {
  loadingTasks.value = true;
  errorMsg.value = "";

  try {
    const token = getToken();

    // Token nahi mila
    if (!token) {
      errorMsg.value = "Please login first.";
      return;
    }

    const response = await fetch(
      "http://localhost:8080/api/tasks",
      {
        method: "GET",
        headers: {
          "Authorization": `Bearer ${token}`,
          "Content-Type": "application/json",
        },
      }
    );

    const data = await response.json();

    // console.log("GET TASKS RESPONSE:", data);

    if (!response.ok) {
      throw new Error(
        data.error || "Failed to load tasks"
      );
    }

    // ⭐ Backend se aaye tasks ko Vue mein save karo
    tasks.value = data;

  } catch (error) {
    console.error("GET TASKS ERROR:", error);

    errorMsg.value =
      error.message || "Couldn't load tasks.";
  } finally {
    loadingTasks.value = false;
  }
}

// =========================
// ADD TASK
// =========================

async function addTask() {
  errorMsg.value = "";

  if (!title.value.trim()) {
    errorMsg.value =
      "Give the task a title before adding it.";
    return;
  }

  try {
    adding.value = true;

    const token = getToken();

    if (!token) {
      errorMsg.value = "Please login first.";
      return;
    }

    const response = await fetch(
      "http://localhost:8080/api/tasks",
      {
        method: "POST",

        headers: {
          "Content-Type": "application/json",

          // ⭐ JWT
          "Authorization": `Bearer ${token}`,
        },

        body: JSON.stringify({
          title: title.value,
          description: description.value,
          complete: false,
        }),
      }
    );

    const data = await response.json();

    // console.log("CREATE TASK RESPONSE:", data);

    if (!response.ok) {
      throw new Error(
        data.error || "Failed to add task"
      );
    }

    // ⭐ New task list ke beginning mein
    tasks.value.unshift(data);

    // Form clear
    title.value = "";
    description.value = "";

  } catch (error) {
    console.error("CREATE TASK ERROR:", error);

    errorMsg.value =
      error.message ||
      "Couldn't add that task. Try again.";

  } finally {
    adding.value = false;
  }
}

// =========================
// DELETE TASK
// =========================

async function deleteTask(id) {
  errorMsg.value = "";

  try {
    const token = getToken();

    if (!token) {
      errorMsg.value = "Please login first.";
      return;
    }

    const response = await fetch(
      `http://localhost:8080/api/tasks/${id}`,
      {
        method: "DELETE",

        headers: {
          "Authorization": `Bearer ${token}`,
        },
      }
    );

    const data = await response.json();

    // console.log("DELETE RESPONSE:", data);

    if (!response.ok) {
      throw new Error(
        data.error || "Failed to delete task"
      );
    }

    // ⭐ Deleted task ko frontend list se remove karo
    tasks.value = tasks.value.filter(
      (task) => task.id !== id
    );

  } catch (error) {
    // console.error("DELETE TASK ERROR:", error);

    errorMsg.value =
      error.message ||
      "Couldn't delete that task. Try again.";
  }
}

// =========================
// TOGGLE TASK
// =========================

async function toggleTask(task) {
  errorMsg.value = "";

  try {
    const token = getToken();

    if (!token) {
      errorMsg.value = "Please login first.";
      return;
    }

    const newCompleteValue = !task.complete;

    const response = await fetch(
      `http://localhost:8080/api/tasks/${task.id}`,
      {
        method: "PUT",

        headers: {
          "Content-Type": "application/json",

          // ⭐ JWT
          "Authorization": `Bearer ${token}`,
        },

        body: JSON.stringify({
          title: task.title,
          description: task.description,
          complete: newCompleteValue,
        }),
      }
    );

    const data = await response.json();

    // console.log("UPDATE TASK RESPONSE:", data);

    if (!response.ok) {
      throw new Error(
        data.error || "Failed to update task"
      );
    }

    // ⭐ Backend successful hone ke baad state update
    task.complete = newCompleteValue;

  } catch (error) {
    console.error("UPDATE TASK ERROR:", error);

    errorMsg.value =
      error.message ||
      "Couldn't update that task. Try again.";
  }
}
const router=useRouter();
function logout(){
  localStorage.removeItem("token");
  localStorage.removeItem("user");
  router.push('/login')
}

// =========================
// PAGE LOAD
// =========================

onMounted(() => {
  getTasks();
});
</script>
<template>
  <div class="page">

    <!-- Header -->
    <header class="topbar">
      <div class="topbar-inner">
        <div class="brand-mark">
          <span class="brand-dot"></span>
          Tally
        </div>

        <div class="header-actions">
          <div
            class="tally-counter"
            role="status"
            aria-live="polite"
          >
            <svg
              class="ring"
              width="52"
              height="52"
              viewBox="0 0 52 52"
              aria-hidden="true"
            >
              <circle
                cx="26"
                cy="26"
                r="22"
                class="ring-track"
              />

              <circle
                cx="26"
                cy="26"
                r="22"
                class="ring-fill"
                :stroke-dasharray="ringCircumference"
                :stroke-dashoffset="ringOffset"
              />
            </svg>

            <div class="tally-numbers">
              <span class="tally-count">
                {{ completed }}
                <span class="tally-of">
                  /{{ total }}
                </span>
              </span>

              <span class="tally-label">
                done
              </span>
            </div>
          </div>

          <button
            type="button"
            class="logout-btn"
            @click="logout"
          >
            Logout
          </button>
        </div>
      </div>
    </header>

    <main class="content">

      <p
        v-if="errorMsg"
        class="banner-error"
        role="alert"
      >
        {{ errorMsg }}
      </p>

      <div class="layout">

        <!-- Add task -->
        <section class="panel add-panel">

          <h2>Add a task</h2>

          <p class="panel-sub">
            What's next on the list?
          </p>

          <form @submit.prevent="addTask">

            <div class="field">

              <label for="title">
                Title
              </label>

              <input
                id="title"
                v-model="title"
                type="text"
                placeholder="e.g. Reply to client emails"
              />

            </div>

            <div class="field">

              <label for="description">
                Details
              </label>

              <textarea
                id="description"
                v-model="description"
                placeholder="Optional notes..."
              ></textarea>

            </div>

            <button
              type="submit"
              class="btn-primary"
              :disabled="adding"
            >
              <span
                v-if="adding"
                class="spinner"
                aria-hidden="true"
              ></span>

              {{ adding ? "Adding..." : "Add task" }}
            </button>

          </form>

        </section>

        <!-- Task list -->
        <section class="panel list-panel">

          <div class="list-heading">

            <h2>My tasks</h2>

            <span
              v-if="total"
              class="list-count"
            >
              {{ total }} total
            </span>

          </div>

          <!-- Loading -->
          <div
            v-if="loadingTasks"
            class="skeleton-list"
          >

            <div
              class="skeleton-card"
              v-for="n in 3"
              :key="n"
            ></div>

          </div>

          <!-- Empty -->
          <div
            v-else-if="tasks.length === 0"
            class="empty-state"
          >

            <svg
              width="40"
              height="40"
              viewBox="0 0 40 40"
              fill="none"
              aria-hidden="true"
            >
              <rect
                x="4"
                y="4"
                width="32"
                height="32"
                rx="10"
                fill="var(--accent-soft)"
              />

              <path
                d="M14 20L18 24L27 15"
                stroke="var(--accent)"
                stroke-width="2.5"
                stroke-linecap="round"
                stroke-linejoin="round"
              />
            </svg>

            <p>
              <strong>No tasks yet.</strong>
              <br />
              Add your first one on the left.
            </p>

          </div>

          <!-- Tasks -->
          <ul
            v-else
            class="task-list"
          >

            <li
              v-for="task in tasks"
              :key="task.id"
              class="task-card"
              :class="{ 'is-complete': task.complete }"
            >

              <button
                type="button"
                class="check"
                :class="{ checked: task.complete }"
                @click="toggleTask(task)"
                :aria-pressed="task.complete"
                :aria-label="
                  task.complete
                    ? 'Mark as pending'
                    : 'Mark as complete'
                "
              >

                <svg
                  viewBox="0 0 16 16"
                  width="12"
                  height="12"
                  aria-hidden="true"
                >
                  <path
                    d="M3 8.5L6.2 11.5L13 4.5"
                    fill="none"
                    stroke="#fff"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  />
                </svg>

              </button>

              <div class="task-body">

                <h3>
                  {{ task.title }}
                </h3>

                <p v-if="task.description">
                  {{ task.description }}
                </p>

                <span
                  class="status-pill"
                  :class="
                    task.complete
                      ? 'is-done'
                      : 'is-pending'
                  "
                >
                  {{
                    task.complete
                      ? "Completed"
                      : "Pending"
                  }}
                </span>

              </div>

              <button
                type="button"
                class="icon-btn"
                @click="deleteTask(task.id)"
                aria-label="Delete task"
              >

                <svg
                  viewBox="0 0 20 20"
                  width="16"
                  height="16"
                  aria-hidden="true"
                >
                  <path
                    d="M5 6h10M8 6V4.5A1.5 1.5 0 0 1 9.5 3h1A1.5 1.5 0 0 1 12 4.5V6m-6 0 .6 9.4A1.5 1.5 0 0 0 8.1 17h3.8a1.5 1.5 0 0 0 1.5-1.6L14 6"
                    stroke="currentColor"
                    stroke-width="1.6"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    fill="none"
                  />
                </svg>

              </button>

            </li>

          </ul>

        </section>

      </div>

    </main>

  </div>
</template>

<style scoped>
.page {
  min-height: 100vh;
}

.topbar {
  position: sticky;
  top: 0;
  z-index: 10;
  background: rgba(247, 247, 244, 0.85);
  backdrop-filter: blur(10px);
  border-bottom: 1px solid var(--border);
}

.topbar-inner {
  max-width: 1040px;
  margin: 0 auto;
  padding: 16px 28px;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.brand-mark {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  font-family: var(--font-display);
  font-weight: 700;
  font-size: 19px;
  color: var(--ink);
}

.header-actions {
  display: flex;
  align-items: center;
  gap: 16px;
}

.logout-btn {
  border: 1px solid var(--border);
  background: var(--surface);
  color: var(--ink);
  border-radius: var(--radius-sm);
  padding: 9px 14px;
  font-size: 13px;
  font-weight: 600;
  cursor: pointer;
}

.brand-dot {
  width: 9px;
  height: 9px;
  border-radius: 50%;
  background: var(--accent);
  box-shadow: 0 0 0 4px var(--accent-soft);
}

.tally-counter {
  display: flex;
  align-items: center;
  gap: 12px;
}

.ring {
  transform: rotate(-90deg);
}

.ring-track {
  fill: none;
  stroke: var(--border);
  stroke-width: 4;
}

.ring-fill {
  fill: none;
  stroke: var(--accent);
  stroke-width: 4;
  stroke-linecap: round;
  transition: stroke-dashoffset 0.4s ease;
}

.tally-numbers {
  display: flex;
  flex-direction: column;
  line-height: 1.1;
}

.tally-count {
  font-family: var(--font-display);
  font-weight: 700;
  font-size: 16px;
  color: var(--ink);
}

.tally-of {
  font-weight: 500;
  color: var(--ink-faint);
}

.tally-label {
  font-size: 11px;
  color: var(--ink-faint);
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.content {
  max-width: 1040px;
  margin: 0 auto;
  padding: 32px 28px 64px;
}

.banner-error {
  background: var(--danger-soft);
  color: var(--danger);
  font-size: 14px;
  padding: 12px 16px;
  border-radius: var(--radius-sm);
  margin: 0 0 20px;
}

.layout {
  display: grid;
  grid-template-columns: 320px 1fr;
  gap: 24px;
  align-items: start;
}

.panel {
  background: var(--surface);
  border: 1px solid var(--border);
  border-radius: var(--radius-lg);
  padding: 24px;
}

.add-panel {
  position: sticky;
  top: 92px;
}

.add-panel h2,
.list-panel h2 {
  font-size: 18px;
  font-weight: 600;
}

.panel-sub {
  color: var(--ink-soft);
  font-size: 14px;
  margin: 6px 0 20px;
}

.field {
  margin-bottom: 16px;
}

.field label {
  display: block;
  font-size: 13px;
  font-weight: 600;
  color: var(--ink-soft);
  margin-bottom: 7px;
}

.field input,
.field textarea {
  width: 100%;
  box-sizing: border-box;
  padding: 11px 13px;
  font-size: 14.5px;
  font-family: var(--font-body);
  color: var(--ink);
  background: var(--bg);
  border: 1.5px solid var(--border);
  border-radius: var(--radius-sm);
  transition:
    border-color 0.15s ease,
    box-shadow 0.15s ease;
  resize: vertical;
}

.field textarea {
  min-height: 84px;
}

.field input::placeholder,
.field textarea::placeholder {
  color: var(--ink-faint);
}

.field input:focus,
.field textarea:focus {
  outline: none;
  border-color: var(--accent);
  box-shadow: 0 0 0 4px var(--accent-soft);
  background: var(--surface);
}

.btn-primary {
  width: 100%;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  padding: 12px 18px;
  font-size: 14.5px;
  font-weight: 600;
  color: #fff;
  background: var(--accent);
  border: none;
  border-radius: var(--radius-sm);
  cursor: pointer;
  transition:
    background 0.15s ease,
    transform 0.1s ease;
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
  width: 14px;
  height: 14px;
  border: 2px solid rgba(255, 255, 255, 0.4);
  border-top-color: #fff;
  border-radius: 50%;
  animation: spin 0.7s linear infinite;
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}

.list-heading {
  display: flex;
  align-items: baseline;
  justify-content: space-between;
  margin-bottom: 20px;
}

.list-count {
  font-size: 13px;
  color: var(--ink-faint);
}

.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  gap: 14px;
  padding: 56px 20px;
  color: var(--ink-soft);
}

.empty-state p {
  margin: 0;
  font-size: 14.5px;
  line-height: 1.5;
}

.skeleton-list {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.skeleton-card {
  height: 78px;
  border-radius: var(--radius-md);
  background: linear-gradient(
    90deg,
    var(--bg) 25%,
    var(--border) 37%,
    var(--bg) 63%
  );
  background-size: 400% 100%;
  animation: shimmer 1.4s ease infinite;
}

@keyframes shimmer {
  0% {
    background-position: 100% 0;
  }

  100% {
    background-position: 0 0;
  }
}

.task-list {
  list-style: none;
  margin: 0;
  padding: 0;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.task-card {
  display: flex;
  align-items: flex-start;
  gap: 14px;
  padding: 16px;
  border: 1px solid var(--border);
  border-left: 3px solid var(--warning);
  border-radius: var(--radius-md);
  transition:
    border-color 0.15s ease,
    background 0.15s ease;
}

.task-card.is-complete {
  border-left-color: var(--success);
  background: var(--success-soft);
}

.check {
  flex: none;
  width: 24px;
  height: 24px;
  margin-top: 2px;
  border-radius: 50%;
  border: 2px solid var(--ink-faint);
  background: transparent;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.15s ease;
}

.check:hover {
  border-color: var(--accent);
}

.check.checked {
  background: var(--success);
  border-color: var(--success);
}

.check svg path {
  opacity: 0;
  transition: opacity 0.15s ease;
}

.check.checked svg path {
  opacity: 1;
}

.task-body {
  flex: 1;
  min-width: 0;
}

.task-body h3 {
  font-size: 15.5px;
  font-weight: 600;
  margin: 0 0 4px;
  word-break: break-word;
}

.is-complete .task-body h3 {
  text-decoration: line-through;
  color: var(--ink-soft);
}

.task-body p {
  font-size: 13.5px;
  color: var(--ink-soft);
  margin: 0 0 10px;
  line-height: 1.5;
  word-break: break-word;
}

.status-pill {
  display: inline-block;
  font-size: 11.5px;
  font-weight: 600;
  padding: 3px 9px;
  border-radius: 999px;
  text-transform: uppercase;
  letter-spacing: 0.03em;
}

.status-pill.is-pending {
  background: var(--warning-soft);
  color: var(--warning);
}

.status-pill.is-done {
  background: var(--success-soft);
  color: var(--success);
}

.icon-btn {
  flex: none;
  width: 32px;
  height: 32px;
  border-radius: var(--radius-sm);
  border: none;
  background: transparent;
  color: var(--ink-faint);
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition:
    background 0.15s ease,
    color 0.15s ease;
}

.icon-btn:hover {
  background: var(--danger-soft);
  color: var(--danger);
}

@media (max-width: 768px) {
  .topbar-inner {
    padding: 14px 18px;
  }

  .content {
    padding: 20px 16px 48px;
  }

  .layout {
    grid-template-columns: 1fr;
  }

  .add-panel {
    position: static;
  }

  .panel {
    padding: 20px;
    border-radius: var(--radius-md);
  }
}
</style>
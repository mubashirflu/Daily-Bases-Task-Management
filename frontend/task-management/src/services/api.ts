// const BASE_API = "http://localhost:8080/api";
// export async function registerUser(userData: any) {
//   const response = await fetch(`${BASE_API}/register`, {
//     method: "POST",
//     headers: {
//       "Content-Type": "application/json",
//     },
//     body: JSON.stringify(userData),
//   });

//   const data = await response.json();

//   if (!response.ok) {
//     throw new Error(data.error || "Registration failed");
//   }

//   return data;
// }
// export async function login(userdata:any) {
//     const response=await fetch(`${BASE_API}/login`,{
//         method:"POST",
//         headers:{
//             "Content-Type":"application/json"
//         },
//         body:JSON.stringify(userdata)
//     }

//     );
//     return response.json()
// }
// export async function GetTask() {
//     const response=await fetch(`${BASE_API}/tasks`)
//     return await response.json()
// }
const BASE_API = "http://localhost:8080/api";

export async function registerUser(userData: any) {
  const response = await fetch(`${BASE_API}/register`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(userData),
  });

  const data = await response.json();

  if (!response.ok) {
    throw new Error(data.error || "Registration failed");
  }

  return data;
}

export async function login(userData: any) {
  const response = await fetch(`${BASE_API}/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(userData),
  });

  const data = await response.json();

  if (!response.ok) {
    throw new Error(data.error || "Login failed");
  }

  // JWT token save karo
  localStorage.setItem("token", data.token);

  return data;
}

export async function GetTask() {

  // LocalStorage se JWT lena
  const token = localStorage.getItem("token");

  const response = await fetch(`${BASE_API}/tasks`, {
    method: "GET",

    headers: {
      "Authorization": `Bearer ${token}`,
      "Content-Type": "application/json",
    },
  });

  const data = await response.json();

  if (!response.ok) {
    throw new Error(data.error || "Failed to get tasks");
  }

  return data;
}
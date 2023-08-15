import { React, useState } from "react";
import axios from "axios";

function Login() {
  const [formData, setFormData] = useState({
    email: "",
    password: "",
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const options = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      data: JSON.stringify(formData),
      url: "/login",
    }

    try {
      await axios(options);
      console.log("Login successful");
    } catch (err) {
      console.log(err.response.data);
    }
  };

  return (
    <div class="flex flex-col justify-center items-center h-screen">
      <div class="flex flex-col justify-center items-center w-1/2">
        <h1 class="text-[#cdd6f4] text-3xl font-sans font-medium">Welcome back</h1>
        <form class="mt-4 w-4/6" onSubmit={handleSubmit}>
          <div class="w-full">
            <label class="block text-[#eff1f5] text-m font-sans mb-2">Email</label>
            <input class="shadow appearance-none border rounded w-full py-3 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline bg-[#eff1f5]" id="email" type="email" placeholder="Enter your email" name="email" value={formData.email} onChange={handleChange} />
          </div>
          <div class="w-full mt-4">
            <label class="block text-[#eff1f5] text-m font-sans mb-2">Password</label>
            <input class="shadow appearance-none border rounded w-full py-3 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline bg-[#eff1f5]" id="password" type="password" placeholder="Enter your password" name="password" value={formData.password} onChange={handleChange} />
          </div>
          <button class="mt-6 w-full bg-[#cdd6f4] hover:bg-[#eff1f5] text-[#1e1e2e] font-sans font-medium py-4 px-3 rounded focus:outline-none focus:shadow-outline" type="submit">Continue</button>
        </form>
        <p class="flex justify-center mt-4 text-[#eff1f5] font-sans">Don't have an account?<a href="/register" class="ml-1 text-[#cdd6f4]">Sign up</a></p>
        <p class="flex justify-center mt-4 text-[#eff1f5] font-sans font-medium">OR</p>
      </div>
    </div>
  )
}

export default Login;

import { React, useState } from "react";
import axios from "axios";
import Input from '../components/form/Input'
import Button from '../components/form/Button'

function Register() {
  const [formData, setFormData] = useState({
    username: "",
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
      url: "/register",
    }

    try {
      await axios(options);
      console.log("Registration successful");
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <div class="flex flex-col justify-center items-center h-screen">
      <div class="flex flex-col justify-center items-center w-1/2">
        <h1 class="text-[#cdd6f4] text-3xl font-sans font-medium">Create your account</h1>
        <form class="mt-4 w-4/6" onSubmit={handleSubmit}>
          <div class="w-full">
            <Input id="username" label="Profile name" type="text" placeholder="Enter a profile name" name="username" value={formData.username} handleChange={handleChange} />
          </div>
          <div class="w-full mt-4">
            <Input id="email" label="Email" type="email" placeholder="Enter your email" name="email" value={formData.email} handleChange={handleChange} />
          </div>
          <div class="w-full mt-4">
            <Input id="password" label="Password" type="password" placeholder="Enter a password" name="password" value={formData.password} handleChange={handleChange} />
          </div>
          <Button type="submit" text="Continue" />
        </form>
        <p class="flex justify-center mt-4 text-[#eff1f5] font-sans">Already have an account?<a href="/login" class="ml-1 text-[#cdd6f4]">Login</a></p>
        <p class="flex justify-center mt-4 text-[#eff1f5] font-sans font-medium">OR</p>
      </div>
    </div>
  )
}

export default Register;

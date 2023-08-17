import { React, useState } from "react";
import axios from "axios";
import Input from '../components/form/Input'
import Button from '../components/form/Button'
import Alert from '../components/ui/Alert'

function Login() {
  const [alertType, setAlertType] = useState(null);
  const [alertMsg, setAlertMsg] = useState(null);
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
    setAlertMsg(null);
    setAlertType(null);

    const options = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      data: JSON.stringify(formData),
      url: "/login",
    }

    try {
      await axios(options);
      setAlertType("success");
      setAlertMsg("Successful login!");
    } catch (err) {
      setAlertType("error");
      setAlertMsg(err.response.data);
      console.log(err);
    }
  };

  return (
    <div class="flex flex-col items-center h-screen">
      {alertMsg && <Alert type={alertType} text={alertMsg} />}
      <div class="flex flex-col justify-center items-center h-screen w-full">
        <div class="flex flex-col justify-center items-center w-full">
          <h1 class="text-[#cdd6f4] text-3xl font-sans font-medium">Welcome back</h1>
          <form class="flex flex-col items-center mt-4 w-4/6" onSubmit={handleSubmit}>
            <div class="flex flex-col items-center w-full">
              <Input id="email" label="Email" type="email" placeholder="Enter your email" name="email" value={formData.email} handleChange={handleChange} />
            </div>
            <div class="flex flex-col items-center w-full mt-4">
              <Input id="password" label="Password" type="password" placeholder="Enter your password" name="password" value={formData.password} handleChange={handleChange} />
            </div>
            <Button type="submit" text="Continue" />
          </form>
          <p class="flex justify-center mt-4 text-[#eff1f5] font-sans">Don't have an account?<a href="/register" class="ml-1 text-[#cdd6f4]">Sign up</a></p>
          <p class="flex justify-center mt-4 text-[#eff1f5] font-sans font-medium">OR</p>
        </div>
      </div>
    </div>
  )
}

export default Login;

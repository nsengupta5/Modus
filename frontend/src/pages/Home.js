import React, { useState, useEffect } from 'react';
import axios from 'axios';

function Home() {

  const [message, setMessage] = useState('Loading...');

  useEffect(() => {
    axios.get('/intro')
      .then((response) => {
        setMessage(response.data);
        console.log(response);
      })
      .catch((error) => {
        console.error(error);
      })
  }, []);

  return (
    <div class="flex justify-center items-center h-screen">
      <h1 class="text-[#cdd6f4] text-5xl font-sans font-medium">{message}</h1>
    </div>
  )
}

export default Home;

import React, { useState, useEffect } from 'react';

function Alert(props) {
  const { type, text } = props;
  const [show, setShow] = useState(false);

  useEffect(() => {
    const showTimeout = setTimeout(() => {
      setShow(true);
    }, 100);

    const timeout = setTimeout(() => {
      setShow(false);
    }, 3000);

    return () => {
      clearTimeout(showTimeout);
      clearTimeout(timeout);
    }
  }, []);

  const alertType = {
    "success": "bg-[#a6e3a1]", 
    "error": "bg-[#f38ba8]",
    "warning": "bg-[#f9e2af]",
  }

  const alertStyle = `${show ? 'opacity-100 translate-y-0' : 'opacity-0 translate-y-2'} ${alertType[type]} transition-opacity duration-500 ease-in-out flex justify-center border border-red-400 text-red-700 px-4 py-3 rounded relative w-4/6`

  return (
    <div class={alertStyle} role="alert">
      <span class="block sm:inline">{text}</span>
      <span class="absolute top-0 bottom-0 right-0 px-4 py-3">
        <svg class="fill-current h-6 w-6 text-red-500" role="button" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><title>Close</title><path d="M14.348 14.849a1.2 1.2 0 0 1-1.697 0L10 11.819l-2.651 3.029a1.2 1.2 0 1 1-1.697-1.697l2.758-3.15-2.759-3.152a1.2 1.2 0 1 1 1.697-1.697L10 8.183l2.651-3.031a1.2 1.2 0 1 1 1.697 1.697l-2.758 3.152 2.758 3.15a1.2 1.2 0 0 1 0 1.698z"/></svg>
      </span>
    </div>
  )
}

export default Alert;

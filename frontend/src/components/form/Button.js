import React from 'react';

function Button(props) {
  const { type, text, handleClick } = props;
  return (
          <button class="mt-6 w-4/6 bg-[#cdd6f4] hover:bg-[#eff1f5] text-[#1e1e2e] font-sans font-medium py-4 px-3 rounded focus:outline-none focus:shadow-outline" type={type}>{text}</button>
  )
}

export default Button;

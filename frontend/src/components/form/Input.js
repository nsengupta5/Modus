import React from 'react';

function Input(props) {
  const { id, label, type, placeholder, name, value, handleChange } = props;

  return (
    <div class='w-4/6'>
      <label class="block text-[#eff1f5] text-m font-sans mb-2">{label}</label>
      <input class="shadow appearance-none border rounded w-full py-3 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline bg-[#eff1f5]" id={id} type={type} placeholder={placeholder} name={name} value={value} onChange={handleChange} />
    </div>
  )
}

export default Input;

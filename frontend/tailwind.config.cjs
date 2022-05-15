const colors = require('tailwindcss/colors');

const config = {
    content: [
      "./src/**/*.{html,js,svelte,ts}",
    ],
  
    theme: {
      extend: {
          colors: {
              
          } 
      },
    },
  
    plugins: [
        require("daisyui"),
    ],
  };
  
  module.exports = config;
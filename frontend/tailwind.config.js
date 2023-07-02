/** @type {import('tailwindcss').Config} */
module.exports = {
	future: false,
	content: [
		'./app/**/*.{js,ts,jsx,tsx,mdx}',
		'./widgets/**/*.{jx,ts,jsx,tsx,mdx}',
		'./shared/**/*.{jx,ts,jsx,tsx,mdx}',
		'./entities/**/*.{jx,ts,jsx,tsx,mdx}',
		'./components/**/*.{jx,ts,jsx,tsx,mdx}',
		'./utils/**/*.{jx,ts,jsx,tsx,mdx}',
		'./constants/**/*.{jx,ts,jsx,tsx,mdx}',
	],
	theme: {
		extend: {
			colors: {
				primary: '#1BB234',
				'primary-light': '#DFFFE2',
				secondary: '#71C7B7',
				tertiary: '#0093E6',
				accent: '#E18C44',
			},
			maxWidth: {
				'container-xl': 1832,
				'container-md': 1592,
				'container-sm': 1232,
			},
		},
	},
	plugins: [require('tailwind-scrollbar')],
};

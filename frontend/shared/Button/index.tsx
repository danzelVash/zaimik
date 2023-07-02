import { ReactNode } from 'react'

interface IButtonProps {
	children: ReactNode;
	type?: 'button' | 'submit';
	className?: string;
}

const Button: React.FC<IButtonProps> = ({
	children,
	className = '',
	type = 'button',
}) => {
	return (
		<button
			className={`
				${className.includes('bg-') ? '' : 'bg-tertiary hover:bg-[#007cc3]'}
				${className.includes('font-') ? '' : 'font-medium'}
				text-white  lg:text-[28px] md:text-[24px] text-[18px] text-center px-3 lg:py-4 py-3 rounded-3xl w-full transition-colors duration-300
				${className}
			`}
			type={type}
		>
			{children}
		</button>
	);
};

export default Button;

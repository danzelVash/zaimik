import { ReactNode } from 'react';

interface ICalculatorTitleProps {
	children: ReactNode;
}

const CalculatorTitle: React.FC<ICalculatorTitleProps> = ({ children }) => {
	return (
		<div className='font-medium md:text-[30px] text-[22px] leading-tight text-center'>
			{children}
		</div>
	);
};

export default CalculatorTitle;

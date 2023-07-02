import { ReactNode } from 'react'

interface ICalculatorValueProps {
	children: ReactNode;
}

const CalculatorValue: React.FC<ICalculatorValueProps> = ({ children }) => {
	return (
		<div className='font-bold lg:text-[36px] text-[25px] leading-tight uppercase text-center mt-1'>
			{children}
		</div>
	);
};

export default CalculatorValue;

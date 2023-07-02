import { ReactNode } from 'react';

interface ISearchTextProps {
	children: ReactNode;
	className?: string;
}

const SearchText: React.FC<ISearchTextProps> = ({
	children,
	className = '',
}) => {
	return (
		<p
			className={`lg:text-[22px] md:text-[20px] text-[14px] leading-tight font-semibold text-center text-white ${className}`}
		>
			{children}
		</p>
	);
};

export default SearchText;

import { ReactNode } from 'react';

interface ISearchTitleProps {
	children: ReactNode;
}

const SearchTitle: React.FC<ISearchTitleProps> = ({ children }) => {
	return (
		<div className='text-center text-white lg:text-[36px] md:text-[30px] text-[24px] leading-tight font-bold'>
			{children}
		</div>
	);
};

export default SearchTitle;

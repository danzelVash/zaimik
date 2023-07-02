import { ReactNode } from 'react';

interface ISearchWrapperProps {
	children: ReactNode;
}

const SearchWrapper: React.FC<ISearchWrapperProps> = ({ children }) => {
	return (
		<div className='main-gradient relative xl:py-14 lg:py-10 md:py-8 py-6 xl:px-16 md:px-10 px-4 lg:rounded-[150px] rounded-0 shadow-[0px_5px_10px_rgba(0,0,0,.25)]'>
			{children}
		</div>
	);
};

export default SearchWrapper;

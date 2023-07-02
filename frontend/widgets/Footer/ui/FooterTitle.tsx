import { ReactNode } from 'react';

interface IFooterTitleProps {
	children: ReactNode;
}

const FooterTitle: React.FC<IFooterTitleProps> = ({ children }) => {
	return (
		<div className='md:font-bold font-semibold xl:text-[26px] md:text-[32px] text-[20px] leading-tight'>{children}</div>
	)
}

export default FooterTitle
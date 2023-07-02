import { ReactNode } from 'react';

interface ITitleProps {
	children: ReactNode;
	type?: 'page' | 'section';
	styles?: string;
}

const Title: React.FC<ITitleProps> = ({ children, styles = '', type = 'section' }) => {
	const classes: string =
		`font-extrabold lg:text-[60px] md:text-[40px] text-[32px] leading-tight md:text-center text-left ${styles}`;

	return (
		<>
			{type === 'page' ? (
				<h1 className={classes}>{children}</h1>
			) : (
				<h2 className={classes}>{children}</h2>
			)}
		</>
	);
};

export default Title;

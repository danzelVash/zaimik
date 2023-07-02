'use client';

import Image from 'next/image';
import { useEffect, useState } from 'react';

// Move to the top of the page
const scrollToTop = (): void =>
	window.scrollTo({
		top: 0,
		behavior: 'smooth',
	});

const ToTop: React.FC = () => {
	const [isVisible, setIsVisible] = useState<Boolean>(false);

	// Watch scroll and calculate scrolled percent
	const handleScroll = (): void => {
		const scrollPercentage: number = Math.round(
			(window.scrollY /
				(document.documentElement.scrollHeight - window.innerHeight)) *
				100
		);

		if (scrollPercentage >= 25) setIsVisible(true);
		else setIsVisible(false);
	};

	useEffect(() => {
		window.addEventListener('scroll', handleScroll);
		return () => window.removeEventListener('scroll', handleScroll);
	}, []);

	return (
		<div className='lock-padding fixed lg:bottom-[40px] md:bottom-[30px] bottom-[15px] lg:right-[40px] md:right-[30px] right-[15px] z-[15] cursor-pointer' onClick={scrollToTop}>
			<div
				className={`
				main-gradient flex items-center justify-center rounded-full shadow-[0px_5px_10px_rgba(0,0,0,.25)]
				md:w-[80px] w-[60px] md:h-[80px] h-[60px] transition-all duration-300
				${isVisible ? 'opacity-100 visible' : 'opacity-0 invisible'} 
			`}
			>
				<Image
					className='md:w-auto w-[28px] md:h-auto h-[32px]'
					src='/icons/to_top.svg'
					width={37}
					height={42}
					alt=''
				/>
			</div>
		</div>
	);
};

export default ToTop;

'use client';

import { useEffect, useState } from 'react';
import HeaderCity from './HeaderCity';
import HeaderLogo from './HeaderLogo';
import HeaderNav from './HeaderNav';
import HeaderUser from './HeaderUser';
import styles from './styles.module.scss';

const Header: React.FC = () => {
	const [isNavOpened, setIsNavOpened] = useState<Boolean>(false);

	const hideNav = (): void => setIsNavOpened(false);

	useEffect((): void => {
		if(isNavOpened) {
			document.body.classList.add('lock');
		} else {
			document.body.classList.remove('lock');
		}
	}, [isNavOpened]);

	return (
		<header
			className={`${styles.header} lock-padding before:bg-gradient-to-r before:from-secondary before:to-primary before:shadow-[0px_5px_10px_rgba(0,0,0,.25)] md:py-2 py-1.5 fixed z-[20] w-full top-0 left-0`}
		>
			<div className='container-xl flex flex-row justify-between items-center gap-4 relative'>
				<HeaderLogo />
				<HeaderNav isOpened={isNavOpened} hideNav={hideNav} />
				<div className='flex flex-row lg:items-start items-center 2xl:gap-x-7 lg:gap-x-6 md:gap-x-5 gap-x-1.5'>
					<HeaderCity />
					<HeaderUser />
					<div
						onClick={() => setIsNavOpened(prev => !prev)}
						className='lg:hidden block md:w-[46px] md:h-[38px] w-[35px] h-[32px] relative z-[2] cursor-pointer transition-hover duration-300 hover:opacity-70'
					>
						<span
							className={`${
								isNavOpened ? '-rotate-45 md:top-[15.5px] top-[13px]' : 'top-0'
							} transition-all duration-200 w-full absolute left-0 h-[6px] bg-[#000000b3] rounded-xl`}
						></span>
						<span
							className={`${
								isNavOpened ? 'opacity-0' : 'opacity-100'
							} transition-all duration-200 w-full absolute top-1/2 -translate-y-1/2 left-0 h-[6px] bg-[#000000b3] rounded-xl`}
						></span>
						<span
							className={`${
								isNavOpened
									? 'rotate-45 md:bottom-[16px] bottom-[13px]'
									: 'bottom-0'
							} transition-all duration-200 w-full absolute left-0 h-[6px] bg-[#000000b3] rounded-xl`}
						></span>
					</div>
				</div>
			</div>
		</header>
	);
};

export default Header;

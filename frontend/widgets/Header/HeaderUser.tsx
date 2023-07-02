import { PopupContext } from '@/context/PopupContext';
import Image from 'next/image';
import React, { useContext } from 'react';

const HeaderUser: React.FC = () => {
	const ctx = useContext(PopupContext);

	return (
		<div
			onClick={() => ctx.setPopup('auth')}
			className='relative z-[2] text-center 2xl:text-[20px] text-[17px] underline cursor-pointer leading-none transition-opacity duration-200 hover:opacity-70'
		>
			<Image
				className='mx-auto 2xl:h-[50px] lg:h-[40px] md:h-[60px] h-[40px]'
				src='/icons/user.svg'
				width={50}
				height={50}
				alt=''
			/>
			<div className='mt-2 lg:block hidden'>Войти</div>
		</div>
	);
};

export default React.memo(HeaderUser);

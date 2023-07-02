'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';

const LoadingCard: React.FC = () => {
	const router = useRouter();
	const [progress, setProgress] = useState<number>(0);

	useEffect(() => {
		const interval = setInterval(() => {
			setProgress(prev => {
				let newProgress: number = prev + Math.floor(Math.random() * 10) + 5;
				if (newProgress >= 100) {
					newProgress = 100;
					clearInterval(interval);
				}
				return newProgress;
			});
		}, 500);
		return () => clearInterval(interval);
	}, [router]);

	useEffect(() => {
		if(progress === 100) 
			router.push('/companies');
	}, [progress, router]);

	return (
		<div className='md:mt-5 mt-4 bg-tertiary md:rounded-[75px] rounded-xl xl:p-10 md:p-7 p-5 text-center text-white'>
			<div className='lg:text-[36px] md:text-[30px] text-[24px] leading-tight font-semibold'>
				Обработано
			</div>
			<div className='lg:text-[90px] md:text-[60px] text-[50px] leading-tight font-black '>
				{progress}%
			</div>
			<div className='lg:mt-4 mt-3 w-full rounded-[75px] lg:h-[75px] md:h-[60px] h-[50px] bg-[#D9D9D9] relative'>
				<div
					style={{
						width: `${progress}%`,
					}}
					className='transition-width duration-300 absolute rounded-[75px] top-0 left-0 h-full main-gradient shadow-[0px_5px_10px_rgba(0,0,0,.25)]'
				></div>
			</div>
			<p className='lg:mt-6 mt-4 lg:text-[26px] md:text-[21px] text-[16px] leading-tight font-semibold'>
				Осталось подождать совсем немного, скоро мы предложим вам подходящие
				варианты
			</p>
		</div>
	);
};

export default LoadingCard;

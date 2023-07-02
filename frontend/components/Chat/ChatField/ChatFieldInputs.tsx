'use client';

import Button from '@/shared/Button';
import { useRouter } from 'next/navigation';

const ChatFieldInputs = () => {
	const router = useRouter();

	return (
		<>
			<div className='grid lg:grid-cols-4 md:grid-cols-2 grid-cols-1 lg:gap-4 md:gap-3 gap-2'>
				<div className='md:h-[70px] h-[50px]'>
					<input
						className='rounded-xl w-full h-full bg-white md:px-5 px-4 md:text-[20px] text-[16px] text-[#414141] placeholder:text-inherit'
						type='text'
						placeholder='ФИО'
					/>
				</div>
				<div className='md:h-[70px] h-[50px]'>
					<input
						className='rounded-xl w-full h-full bg-white md:px-5 px-4 md:text-[20px] text-[16px] text-[#414141] placeholder:text-inherit'
						type='text'
						placeholder='Место работы'
					/>
				</div>
				<div className='md:h-[70px] h-[50px]'>
					<input
						className='rounded-xl w-full h-full bg-white md:px-5 px-4 md:text-[20px] text-[16px] text-[#414141] placeholder:text-inherit'
						type='text'
						placeholder='Сумма выплаты'
					/>
				</div>
				<div className='md:h-[70px] h-[50px]'>
					<input
						className='rounded-xl w-full h-full bg-white md:px-5 px-4 md:text-[20px] text-[16px] text-[#414141] placeholder:text-inherit'
						type='text'
						placeholder='Дата выплаты'
					/>
				</div>
			</div>
			<div className='flex justify-center'>
				<div
					className='lg:max-w-[500px] w-full'
					onClick={() => router.push('/loading')}
				>
					<Button className='lg:mt-7 md:mt-3 mt-2 rounded-xl w-full mx-auto'>
						Отправить
					</Button>
				</div>
			</div>
		</>
	);
};

export default ChatFieldInputs;

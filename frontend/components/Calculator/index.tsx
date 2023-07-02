'use client';

import Button from '@/shared/Button';
import RangeSlider from '@/shared/RangeSlider';
import { getDayWord } from '@/utils/helpers';
import { FormEvent, useState } from 'react';
import CalculatorTitle from './ui/CalculatorTitle';
import CalculatorValue from './ui/CalculatorValue';
import { useRouter } from 'next/navigation';

const Calculator: React.FC = () => {
	const router = useRouter();
	const [sumValue, setSumValue] = useState<[number, number]>([1_000, 1_000]);
	const [timeValue, setTimeValue] = useState<[number, number]>([1, 1]);

	const [, currentSum]: number[] = sumValue;
	const [, currentTime]: number[] = timeValue;

	const dayWord: string = getDayWord(currentTime);

	const formHandler = (event: FormEvent<HTMLFormElement>): void => {
		event.preventDefault();
		router.push('/chat');
	};

	return (
		<div
			id='calculator'
			className='
				xl:px-14 relative mt-6
				before:content-[""] before:absolute before:left-0 before:top-1/2 before:-translate-y-1/2 
				before:w-[36px] before:h-3/4 before:bg-secondary before:rounded-l-3xl xl:before:block before:hidden
				after:content-[""] after:absolute after:right-0 after:top-1/2 after:-translate-y-1/2 
				after:w-[36px] after:h-3/4 after:bg-primary after:rounded-r-3xl xl:after:block after:hidden
			'
		>
			<div className='main-gradient rounded-3xl py-5 lg:px-7 md:px-12 px-5 shadow-[0px_5px_10px_rgba(0,0,0,.25)]'>
				<form onSubmit={formHandler} className='grid lg:grid-cols-2 grid-cols-1 gap-x-10 lg:gap-y-8 md:gap-y-7 gap-y-6 items-center'>
					<div>
						<CalculatorTitle>Сумма займа</CalculatorTitle>
						<CalculatorValue>{currentSum} руб</CalculatorValue>
						<RangeSlider
							value={sumValue}
							setValue={setSumValue}
							min={1_000}
							max={50_000}
							className='mt-1'
						/>
					</div>
					<div>
						<CalculatorTitle>Срок</CalculatorTitle>
						<CalculatorValue>
							{currentTime} {dayWord}
						</CalculatorValue>
						<RangeSlider
							value={timeValue}
							setValue={setTimeValue}
							min={1}
							max={30}
							className='mt-1'
						/>
					</div>
					<div className='text-center lg:text-[22px] md:text-[26px] text-[18px] leading-tight font-bold uppercase md:space-y-4 space-y-2 md:mt-0 mt-2'>
						<div>вы возвращаете: 1255,50 руб.</div>
						<div>дата возврата - 01.01.2024</div>
					</div>
					<div>
						<Button
							className='
								bg-accent hover:bg-[#dc7c2a] shadow-[0px_5px_10px_rgba(0,0,0,.25)] lg:py-7 uppercase font-extrabold
							'
							type='submit'
						>
							Оформить заявку
						</Button>
					</div>
				</form>
				<div className='md:mt-7 mt-6 text-center leading-tight lg:text-[20px] md:text-[26px] text-[18px] font-bold uppercase'>
					Стоимость услуги - 600 руб.
				</div>
			</div>
		</div>
	);
};

export default Calculator;

interface IObtainingStagesLinepProps {
	index: number;
}

const ObtainingStagesLine: React.FC<IObtainingStagesLinepProps> = ({
	index,
}) => {
	const offset: string =
		index % 2
			? 'lg:left-[125px] md:left-[48px] left-[25px] lg:top-auto md:top-[125px] top-[69px]'
			: 'lg:right-[145px] right-auto lg:left-auto md:left-[48px] left-[25px] md:bottom-[131px] bottom-[69px]';
	return (
		<div
			className={`
				main-gradient absolute rounded-3xl lg:bottom-[48px] lg:h-[8px] md:h-[calc(100%_/_2_-_209px)] h-[calc(100%_/_2_-_105px)] lg:w-[calc(100%_/_2_-_219px)] md:w-[8px] w-[4px] ${offset}
			`}
		>
			<span
				className='
					leading-none w-0 h-0 border-transparent md:border-[15px] border-[8px] lg:border-r-0 border-l-primary md:border-l-[30px] border-l-[16px]
					absolute lg:top-[50%] md:top-[calc(100%_-_16px)] top-[calc(100%_-_9px)] lg:right-[-6px] right-1/2 z-[1] lg:-translate-y-1/2 translate-x-1/2 lg:rotate-0 rotate-90
				'
			></span>
		</div>
	);
};

export default ObtainingStagesLine;

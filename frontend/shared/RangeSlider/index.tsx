import Slider from 'react-range-slider-input';
import 'react-range-slider-input/dist/style.css';
import './styles.scss';
import { Dispatch, SetStateAction } from 'react';

interface IRangeSliderProps {
	value: [number, number];
	setValue: Dispatch<SetStateAction<[number, number]>>;
	min: number;
	max: number;
	step?: number;
	className?: string;
}

const RangeSlider: React.FC<IRangeSliderProps> = ({
	min,
	max,
	value,
	setValue,
	className = '',
}) => {
	return (
		<div>
			<div
				className={`
					flex flex-row items-center justify-between gap-4 text-[#DFFFE2] md:text-[22px] text-[18px] font-bold leading-tight uppercase
					${className}
				`}
			>
				<div>{min}</div>
				<div>{max}</div>
			</div>
			<Slider
				className='single-range-slider'
				min={min}
				max={max}
				value={value}
				onInput={setValue}
				thumbsDisabled={[true, false]}
				rangeSlideDisabled={true}
			/>
		</div>
	);
};

export default RangeSlider;

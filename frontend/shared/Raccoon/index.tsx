import Image from 'next/image';

interface IRaccoonProps {
	src: string;
	width: number;
	height: number;
	className?: string;
}

const Raccoon: React.FC<IRaccoonProps> = ({
	src,
	width,
	height,
	className = '',
}) => {
	return (
		<div className={className}>
			<Image src={src} width={width} height={height} alt='' />
		</div>
	);
};

export default Raccoon;
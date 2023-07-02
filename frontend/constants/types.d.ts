export interface ICity {
	readonly id;
	name: string;
}

export interface IGeneral {
	email: string;
}

export interface INavLink {
	readonly id: number;
	path: string;
	name: string;
	icon: {
		src: string;
		size: [number, number];
	};
}

export interface IFooterLink {
	readonly id: number;
	path: string;
	name: string;
}

export interface ILoanCondition {
	readonly id: number;
	icon: {
		src: string;
		size: [number, number];
	},
	title: string;
	text: string;
}

export interface IPaymentMethod {
	readonly id: number;
	src: string;
	size: [number, number];
}

export interface IObtainingStage {
	readonly id: number;
	styles: string;
	name: string;
	icon: {
		src: string;
		size: [number, number];
	}
}
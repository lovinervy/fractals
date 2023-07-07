use num::{Complex, complex::ComplexFloat};
use image::{DynamicImage, Luma};

fn poincare(z: Complex<f64>, g2: Complex<f64>, g3: Complex<f64>) -> Complex<f64> {
    let a = (Complex::powi(&z, 4) + g2 * Complex::powi(&z, 2) / 2.0 + 2.0 * g3 * z + Complex::powi(&z, 2) / 16.0) / (4.0 * Complex::powi(&z, 3) - g2 * z - g3);
    return a;
}

fn fix_display(orientation_point: u32, size: u32) -> f64 {
    let point = orientation_point as f64;
    let size = size as f64;
    return 4.0 * point / (size - 1.0) - 2.0;
}

fn calculate(f: fn(Complex<f64>, Complex<f64>, Complex<f64>) -> Complex<f64>, x1: Complex<f64>, x2: Complex<f64>, size: u32, gray: u8, filename: &str){
    let gray_offset: u8 = 255 / gray;
    let mut set_tone = false;
    let mut img = DynamicImage::new_luma8(size, size).to_luma8();
    for y in 0..=size - 1 {
        let zy = fix_display(y, size);
        for x in 0..=size - 1 {
            let zx = fix_display(x, size);
            let mut z = Complex::new(zx, zy);
            for i in 0..=gray {
                if z.abs() as f64 > 2.0 {
                    img.put_pixel(x, y, Luma([gray_offset * i]));
                    set_tone = true;
                    break;
                }
                z = f(z, x1, x2);
            }
            if set_tone == false {
                img.put_pixel(x, y, Luma([170]));
            }
            set_tone = false;
        }
    img.save(filename).unwrap();
    }
}

fn main() {
    let g2: Complex<f64> = Complex::new(0.0, 0.0);
    let g3: Complex<f64> = Complex::new(1.0, 0.0);

    let size = 1000;
    let iterations: u8 = 32;
    let filename = "example.png";
    calculate(poincare, g2, g3, size, iterations, filename);
}

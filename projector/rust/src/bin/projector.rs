use clap::Parser;

fn main() {
    let options = projector::opts::ProjectorOptions::parse();
    println!("{:?}", options)
}

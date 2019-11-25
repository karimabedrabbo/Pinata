//
//  ColorCollectionViewDelegate.swift
//  PINATA
//
//  Created by Karim Abedrabbo on 11/18/18.
//  Copyright © 2018 Karim Abedrabbo. All rights reserved.
//

import Foundation
import UIKit

class ColorCollectionViewDelegate: NSObject, UICollectionViewDataSource, UICollectionViewDelegate, UICollectionViewDelegateFlowLayout {
    
    weak var colorDelegate : ColorDelegate?
    
    let colors: [UIColor] = {
        let differenceValue: CGFloat =  0.2
        let regularColors: [UIColor] = [.black, .white, .green, .blue, .red, .magenta, .yellow, .cyan, .brown, .purple, .orange]
        
        let darkerAndLighter: [UIColor] = [UIColor.black.withAlphaComponent(0.5), UIColor.white.withAlphaComponent(0.5),UIColor.green.darkerHsb(by: 35), UIColor.blue.lighterHsb(by: 60), UIColor.red.lighterHsb(by: 60), UIColor.magenta.darkerHsb(by: 20), UIColor.yellow.darkerHsb(by: 30), UIColor.cyan.darkerHsb(by: 20), UIColor.brown.lighterHsb(by: 20), UIColor.purple.lighterHsb(by: 30), UIColor.orange.darkerHsb(by: 30)]
            
        


        return regularColors + regularColors + darkerAndLighter + darkerAndLighter
        }()
    
    
    
    override init() {
        super.init()
        
    }
    
    func collectionView(_ collectionView: UICollectionView, numberOfItemsInSection section: Int) -> Int {
        return colors.count
    }
    
    func collectionView(_ collectionView: UICollectionView, didSelectItemAt indexPath: IndexPath) {
        if indexPath.item < 11 || 22 <= indexPath.item && indexPath.item < 33 {
            if colors[indexPath.item] == UIColor.black || colors[indexPath.item] == UIColor.black.withAlphaComponent(0.5){
            
               
                    colorDelegate?.changeColor(textColor: UIColor.white, backgroundColor: colors[indexPath.item])
                
                
            } else if colors[indexPath.item] == UIColor.white || colors[indexPath.item] == UIColor.white.withAlphaComponent(0.5) {
                
                
                    colorDelegate?.changeColor(textColor: UIColor.black, backgroundColor: colors[indexPath.item])
                
                
            }
            else {
                colorDelegate?.changeColor(textColor: UIColor.black, backgroundColor: colors[indexPath.item].withAlphaComponent(0.7))
            }
        }
        else if 11 <= indexPath.item && indexPath.item < 22 || 33 <= indexPath.item && indexPath.item < 44  {
            colorDelegate?.changeColor(textColor: colors[indexPath.item], backgroundColor: UIColor.clear)
            
            
        }
    }
    
    func numberOfSections(in collectionView: UICollectionView) -> Int {
        return 1
    }
    
    func collectionView(_ collectionView: UICollectionView, cellForItemAt indexPath: IndexPath) -> UICollectionViewCell {
        let cell  = collectionView.dequeueReusableCell(withReuseIdentifier: "ColorCollectionViewCell", for: indexPath) as! ColorCollectionViewCell
        
        if indexPath.item < 11 || 22 <= indexPath.item && indexPath.item < 33 {
            cell.colorView.layer.borderColor = UIColor.clear.cgColor
            cell.colorView.backgroundColor = colors[indexPath.item]
            cell.colorView.layer.borderWidth = 0.0
        } else if 11 <= indexPath.item && indexPath.item < 22 || 33 <= indexPath.item && indexPath.item < 44 {
            cell.colorView.layer.borderColor = colors[indexPath.item].cgColor
            cell.colorView.layer.borderWidth = 7.0
            cell.colorView.backgroundColor = UIColor.clear
        }
        
        if cell.colorView.layer.borderColor == UIColor.black.cgColor && cell.colorView.backgroundColor == UIColor.black {
            cell.colorView.layer.borderColor = UIColor.white.cgColor
        }
        
        
        
        
        return cell
    }
    

    
    
}
